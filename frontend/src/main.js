import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import { GetUsername, IsLogged } from '../wailsjs/go/main/App';

IsLogged()
    .then((result) =>  {
        if(result){
            document.querySelector('#app').innerHTML = `
            <img id="logo" class="logo">
            <div class="result" id="result">greet()</div>
            `;
        }else{
            document.querySelector('#app').innerHTML = `
            <img id="logo" class="logo">
            <div class="result" id="result">Login</div>
            <div class="input-box" id="input">
                <input class="input" id="username" type="text" autocomplete="off" />
                <button class="btn" onclick="greet()">Greet</button>
            </div>
            </div>
            `;
        }
        resultElement.innerText = result;
    })

document.getElementById('logo').src = logo;

let nameElement = document.getElementById("name");
nameElement.focus();
let resultElement = document.getElementById("result");

window.greet = function () {
    let name = nameElement.value;

    if (name === "") return;

    try {

        GetUsername()
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
