import './style.css';
import './app.css';

import { FetchPasswords } from '../wailsjs/go/main/App';

document.addEventListener("DOMContentLoaded", () => {
    // Sample JSON data
    const jsonData = '[{"ID":1,"Url":"https://example.com","PasswordHash":"hash1"},{"ID":2,"Url":"https://test.com","PasswordHash":"hash2"},{"ID":3,"Url":"https://dummy.com","PasswordHash":"hash3"}]';
    
    function displayPasswords() {
        FetchPasswords()
            .then((result) =>  {
                const passwords = JSON.parse(result);
                const passwordList = document.getElementById('passwordList');
                passwordList.innerHTML = ''; // Clear existing list items
                
                passwords.forEach(password => {
                    const listItem = document.createElement('li');
                    listItem.innerHTML = `
                        <div class="password-block">
                            <p><strong>URL:</strong> <a href="${password.Url}" target="_blank">${password.Url}</a></p>
                            <p><strong>Password Hash:</strong> ${password.PasswordHash}</p>
                        </div>
                    `;
                    passwordList.appendChild(listItem);
                });
            })
            .catch((err) => {
                console.error(err);
            });
    }

    // Call the function with the JSON data
    displayPasswords();
});


async function fetchPasswords() {
    try {
        const passwords = await FetchPasswords();
        const passwordList = document.getElementById('passwordList');
        const title = document.getElementById('title');
        title.textContent = "TEST"
        passwordList.innerHTML = ''; // Clear existing list items
        passwords.forEach(password => {
            console.log(password)
            const listItem = document.createElement('li');
            listItem.innerHTML = `
                <strong>ID:</strong> ${password.ID}<br>
                <strong>URL:</strong> ${password.Url}<br>
                <strong>Password Hash:</strong> ${password.PasswordHash}
            `;
            passwordList.appendChild(listItem);
        });
    } catch (err) {
        console.error(err); 
    }
};

document.getElementById('actionButton').onclick = fetchPasswords;
