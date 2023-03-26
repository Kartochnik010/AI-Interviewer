const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const sendBtn = document.querySelector('#send')

const username = document.querySelector('#username')
const YearsOfCommercialExperience = document.querySelector('#YearsOfCommercialExperience')
const CurrentPosition = document.querySelector('#CurrentPosition')
const DesiredPosition = document.querySelector('#DesiredPosition')
const stack = document.querySelector('#stack')

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);



send.onclick = () => {
    const message = {
		name: username.value,
        YearsOfCommercialExperience: YearsOfCommercialExperience.value,
        CurrentPosition: CurrentPosition.value,
		content: input.value,
	}
    console.log(message);
    ws.send(JSON.stringify(message));
    input.value = "";
};



/**
 * Insert a message into the UI
 * @param {Message that will be displayed in the UI} messageObj
 */
function insertMessage(messageObj) {
	// Create a div object which will hold the message
	const message = document.createElement('div')

	// Set the attribute of the message div
	message.setAttribute('class', 'chat-message')
	console.log("name: " +messageObj.username + " content: " + messageObj.content)
	message.textContent = `${messageObj.username}: ${messageObj.content}`

	// Append the message to our chat div
	messages.appendChild(message)

	// Insert the message as the first message of our chat
	messages.insertBefore(message, messages.firstChild)
}

const fillForm = (name) => {
    if (name === 'ilyas') {
        username.value = 'Ilyas'
        YearsOfCommercialExperience.value = '1'
        CurrentPosition.value = 'Junior'
        DesiredPosition.value = 'Strong Junior'
        stack.value = 'Golang'
    }
    if (name === 'eldana') {
        username.value = 'Eldana'
        YearsOfCommercialExperience.value = '3'
        CurrentPosition.value = 'Middle'
        DesiredPosition.value = 'Senior'
        stack.value = 'iOS'
    }
}
// ws.onmessage = function (msg) {
//     insertMessage(JSON.parse(msg.data))
// };


const toggleChat = () => {
    const intro = document.querySelector('.intro')
    intro.classList.toggle("thin")
    
    const chat = document.querySelector('.chat')
    chat.classList.toggle("thin")
}

document.querySelector("body").innerHTML+= "<hr>js loaded";