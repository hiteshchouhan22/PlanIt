# Welcome to your Ticket Booking App 🎟️

This is a full-stack ticket booking application built with [React Native](https://reactnative.dev/) and [Golang](https://golang.org/), providing a seamless event discovery and booking experience.

## Get Started

1. **Install Mobile Dependencies**

   Navigate to the Mobile directory and install the required packages:

   ```bash
   cd Mobile
   npm install
   ```
2. **Install Backend Dependencies**

   Navigate to the backend directory and install the necessary Go modules:

   ```bash

   cd backend
   go mod tidy
   ```

3.  **Start the Application**

    You can start the entire application using Docker:

    ```bash
    
    make start
    ```

    Alternatively, you can start the frontend and backend separately:

    💠Mobile:

    ```bash

    cd frontend
    npx expo start
    ```

    💠Backend:

    ```bash

    cd backend
    go run main.go
    ```

    In the output, you'll find options to open the app in a

    - [development build](https://docs.expo.dev/develop/development-builds/introduction/)
    - [Android emulator](https://docs.expo.dev/workflow/android-studio-emulator/)
    - [iOS simulator](https://docs.expo.dev/workflow/ios-simulator/)
    - [Expo Go](https://expo.dev/go), a limited sandbox for trying out app development with Expo

    You can start developing by editing the files inside the **app** directory. This project uses [file-based routing](https://docs.expo.dev/router/introduction).

    ## Learn more

      To learn more about contributing in our project, look at the following resources:

     - [Expo documentation](https://docs.expo.dev/): Learn fundamentals, or go into advanced topics with our [guides](https://docs.expo.dev/guides).
     - [Learn Expo tutorial](https://docs.expo.dev/tutorial/introduction/): Follow a step-by-step tutorial where you'll create a project that runs on Android, iOS, and the web.
     - [Golang Documentation](https://golang.org/doc/):  Learn more about GO From Doc.
     - [GORM Docs](https://gorm.io/docs/index.html): Explore Go’s powerful features for ORM.
  
