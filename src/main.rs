#[macro_use]
extern crate dotenv_codegen;

use actix_web::{get, HttpServer, App, Responder, HttpResponse};
use dotenv::dotenv;

#[get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello world")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();

    HttpServer::new(|| {
        App::new()
            .service(hello)
    })
        .bind(format!(
            "{}:{}",
            dotenv!("ACTIX_HOST"),
            dotenv!("ACTIX_PORT")
        ))?
        .run()
        .await
}