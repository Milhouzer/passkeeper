import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import { IsLogged, Login } from '../wailsjs/go/main/App';

IsLogged()
    .then((result) =>  {
        if(result){
            document.querySelector('#app').innerHTML = `
            <img id="logo" class="logo">
            <div class="result" id="result">greet()</div>
            `;
        }else{
            // document.querySelector('#app').innerHTML = `
            // <img id="logo" class="logo">
            // <div class="result" id="result">Login</div>
            // <div class="input-box" id="input">
            //     <input class="input" id="username" type="text" autocomplete="off" />
            //     <input class="input" id="password" type="text" autocomplete="off" />
            //     <button class="btn" onclick="Login()">Greet</button>
            // </div>
            // </div>
            // `;
            document.querySelector('#app').innerHTML = `
            <!DOCTYPE html>
            <html lang="en">
            <head>
                <meta charset="UTF-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <title>Manage Passwords</title>
                <link rel="stylesheet" href="styles.css">
            </head>
            <body>
                <div class="container">
                    <header>
                        <h1>Manage Passwords</h1>
                    </header>
                    <main>
                        <div class="list-container">
                            <ul class="scrollable-list" id="passwordList">
                                <!-- Example list items -->
                                <li class="list-item">
                                    <span class="name">Name 1</span>
                                    <span class="value">Value 1</span>
                                </li>
                                <li class="list-item">
                                    <span class="name">Name 2</span>
                                    <span class="value">Value 2</span>
                                </li>
                                <li class="list-item">
                                    <span class="name">Name 1</span>
                                    <span class="value">Value 1</span>
                                </li>
                                <li class="list-item">
                                    <span class="name">Name 2</span>
                                    <span class="value">Value 2</span>
                                </li>
                                <!-- Add more list items here -->
                            </ul>
                        </div>
                        <div class="action-container">
                            <input type="text" id="inputField" placeholder="Enter new password">
                            <button id="actionButton">Submit</button>
                        </div>
                    </main>
                </div>

                <script>
                    document.getElementById('actionButton').addEventListener('click', function() {
                        const inputValue = document.getElementById('inputField').value;
                        if (inputValue) {
                            // Call the function with the input value
                            document.getElementById('inputField').value = '';
                        } else {
                            alert('Please enter a value.');
                        }
                    });
                </script>
            </body>
            </html>

            `;
                    

        }
    })

document.getElementById('logo').src = logo;

let usernameElement = document.getElementById("username");
let pwdElement = document.getElementById("password");
let resultElement = document.getElementById("result");

nameElement.focus();

window.greet = function () {
    let name = nameElement.value;

    if (name === "") return;

    try {

        Login(usernameElement.innerText, pwdElement.innerText)
            .then((result) =>  {
                resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

window.login = function() {

}
