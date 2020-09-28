#[macro_use]
extern crate dotenv_codegen;

use dotenv::dotenv;
use teloxide::{dispatching::update_listeners, prelude::*};

use std::{convert::Infallible, env, net::SocketAddr};
use tokio::sync::mpsc;
use warp::Filter;

use reqwest::StatusCode;

#[tokio::main]
async fn main() {
    run().await;
}

async fn run() {
    dotenv().ok();

    teloxide::enable_logging!();
    log::info!("Launching...");

    let bot = Bot::from_env();

    let cloned_bot = bot.clone();
    teloxide::repl_with_listener(
        bot,
        |message| async move {
            message.answer_dice().send().await?;
            ResponseResult::<()>::Ok(())
        },
        webhook(cloned_bot).await,
    ).await;
}

async fn handle_rejection(error: warp::Rejection) -> Result<impl warp::Reply, Infallible> {
    log::error!("Cannot process the request due to: {:?}", error);
    Ok(StatusCode::IM_A_TEAPOT)
}

pub async fn webhook<'a>(bot: Bot) -> impl update_listeners::UpdateListener<Infallible> {
    let host = env::var("HOST").expect("HOST variable is not specified");
    let token = env::var("TELOXIDE_TOKEN").expect("TELOXIDE_TOKEN variable is not specified");
    let url = format!("https://{}/bot{}", host, token);

    bot.set_webhook(url).send().await.expect("Cannot setup a webhook");

    let (tx, rx) = mpsc::unbounded_channel();

    let server = warp::post()
        .and(warp::path(format!("bot{}", dotenv!("TELOXIDE_TOKEN"))))
        .and(warp::body::json())
        .map(move |json: serde_json::Value| {
            let try_parse = match serde_json::from_str(&json.to_string()) {
                Ok(update) => Ok(update),
                Err(error) => {
                    log::error!(
                        "Cannot parse an update\nError: {:?}\nValue: {:?}\n",
                        error,
                        json
                    );

                    Err(error)
                }
            };

            if let Ok(update) = try_parse {
                tx.send(Ok(update)).expect("Cannot send an incoming update from webhook")
            }

            StatusCode::OK
        }).recover(handle_rejection);

    let serve = warp::serve(server);
    let address = format!("0.0.0.0:{}", dotenv!("PORT"));

    tokio::spawn(serve.run(address.parse::<SocketAddr>().unwrap()));
    log::info!("Bot started on {} port", dotenv!("PORT"));

    rx
}