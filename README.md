# gRPC-game

## Overview
gRPC-game is a multiplayer game that uses gRPC for communication between the server and clients. This project demonstrates the use of gRPC in a real-time gaming environment.

## Features
- Real-time multiplayer gameplay
- Efficient communication using gRPC
- Leaderboard
- Scalable server architecture

## Prerequisites
- .NET 6.0 SDK
- Visual Studio or VS Code
- Protobuf compiler (protoc)

## Getting Started
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/gRPC-game.git
    ```
2. Navigate to the project directory:
    ```sh
    cd gRPC-game
    ```
3. Restore the dependencies:
    ```sh
    dotnet restore
    ```
4. Compile the protobuf files:
    ```sh
    protoc -I=Protos --csharp_out=Protos/Generated Protos/*.proto
    ```
5. Run the server:
    ```sh
    dotnet run --project Server
    ```
6. Run the client:
    ```sh
    dotnet run --project Client
    ```

## Usage
- Start the server and then run multiple clients to join the game.
- Follow the on-screen instructions to play the game.

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License.

## Contact
For any questions or feedback, please contact akshat.g1707@gmail.com

