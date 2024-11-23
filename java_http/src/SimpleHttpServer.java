package java_http.src;

import java.io.*;
import java.net.*;
import java.util.concurrent.*;

public class SimpleHttpServer {
    private final int port;
    private final ExecutorService threadPool;

    public SimpleHttpServer(int port, int poolSize) {
        this.port = port;
        this.threadPool = Executors.newFixedThreadPool(poolSize);
    }

    public void start() throws IOException {
        try (ServerSocket serverSocket = new ServerSocket(port)) {
            System.out.println("Server started on port " + port);

            while (true) {
                Socket clientSocket = serverSocket.accept();
                threadPool.execute(new ClientHandler(clientSocket));
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) throws IOException {
        SimpleHttpServer server = new SimpleHttpServer(8080, 10);
        server.start();
    }
}

class ClientHandler implements Runnable {
    private final Socket clientSocket;

    public ClientHandler(Socket socket) {
        this.clientSocket = socket;
    }

    @Override
    public void run() {
        try (BufferedReader in = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
             PrintWriter out = new PrintWriter(clientSocket.getOutputStream(), true)) {

            String requestLine;
            while ((requestLine = in.readLine()) != null) {
                if (requestLine.isEmpty()) {
                    break;
                }
                // System.out.println(requestLine);
            }

            out.println("HTTP/1.1 200 OK");
            out.println("Content-Type: text/plain");
            out.println();
            out.println("Hello, World!r\n");

        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            try {
                clientSocket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}