use futures::stream::StreamExt;

use async_std::net::TcpListener;
use async_std::prelude::*;
// ANCHOR: main_func
use async_std::task::spawn;

#[async_std::main]
async fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").await.unwrap();
    listener
        .incoming()
        .for_each_concurrent(/* limit */ None, |stream| async move {
            let stream = stream.unwrap();
            spawn(handle_connection(stream));
        })
        .await;
}
// ANCHOR_END: main_func

use async_std::io::{Read, Write};

async fn handle_connection(mut stream: impl Read + Write + Unpin) {
    let mut buffer = [0; 1024];
    stream.read(&mut buffer).await.unwrap();
    let status_line = "HTTP/1.1 200 OK\r\n\r\n";
    let contents = "Hello World!\n";
    let response = format!("{status_line}{contents}");
    stream.write(response.as_bytes()).await.unwrap();
    stream.flush().await.unwrap();
}
