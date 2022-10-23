# miniDiscord
Basic, simple, scalable :)

# Example
![miniDiscord-1](https://user-images.githubusercontent.com/75096034/194345156-2f24770d-4d5c-448e-b4f9-65ba8f6205ab.gif)

# Usage
`go run main.go`

### Database
Implement database using `DAO` struct.

Currently using Google Cloud's Firebase.

// Firebase Walkthrough (<https://firebase.google.com/docs/admin/setup#initialize-sdk>)

// SDK Setup (<https://firebase.google.com/docs/admin/setup#go>)
### Setting up Firebase
1. Create a project in Firebase. (<https://console.firebase.google.com/>)
2. Create a service account.
3. Download the service account's private key.
   Go to `Click on your New Project` -> `Project Settings` -> `Service Accounts` -> `Generate new private key`.
4. Rename the private key to `secret.json` and place it in the root directory.

## Libraries Used
- gin
- melody
- go-away
