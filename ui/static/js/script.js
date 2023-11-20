const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const send = document.querySelector('#send')

const username = document.querySelector('#username')
const YearsOfCommercialExperience = document.querySelector('#YearsOfCommercialExperience')
const CurrentPosition = document.querySelector('#CurrentPosition')
const DesiredPosition = document.querySelector('#DesiredPosition')
const stack = document.querySelector('#stack')

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

var converter = new showdown.Converter(),
    text      = '# hello, markdown!',
    html      = converter.makeHtml(text);


const sendChat = () => {
    const userdata = {
		username: document.querySelector('#username').value,
		content: document.querySelector('#textarea').value,
	}
    console.log("sent: "+ userdata.username, userdata.content)
    ws.send("chat," + userdata.username+ ","+ userdata.content);
    document.querySelector('#textarea').value = "";
    insertMessage(userdata.content)
}



/**
 * Insert a message into the UI
 * @param {Message that will be displayed in the UI} messageObj
 */
function insertMessage(input) {
	// Create a div object which will hold the message
	const message = document.createElement('div')

	// Set the attribute of the message div
	message.setAttribute('class', 'chat-message')
	message.textContent = input

	// Append the message to our chat div
	document.querySelector('#messages').appendChild(message)

	// Insert the message as the first message of our chat
	document.querySelector('#messages').insertBefore(message, messages.firstChild)
}


ws.onmessage = function (msg) {
    console.log("Got:"+msg.data);
    insertMessage(msg.data)
};

const startRoadmap = () => {
    let userdata = {
        username: document.querySelector('#username').value,
        YearsOfCommercialExperience: document.querySelector('#YearsOfCommercialExperience').value,
        CurrentPosition: document.querySelector('#CurrentPosition').value,
        DesiredPosition: document.querySelector('#DesiredPosition').value,
        stack: document.querySelector('#stack').value
    }
    
    let err = validateUser(userdata)
    if (err != "") {
        alert(err)
        return
    }

    const intro = document.querySelector('.intro')
    intro.classList.toggle("thin")
    
    const chat = document.querySelector('.chat')
    chat.classList.toggle("thin")

    document.querySelectorAll(".filler").forEach(e => {e.classList.toggle("thin")})

    
    ws.send("start,"+ userdata.username + "," + userdata.YearsOfCommercialExperience + "," + userdata.CurrentPosition + "," + userdata.DesiredPosition + "," + userdata.stack  )

}


const startInterview = () => {
    let userdata = {
        username: document.querySelector('#username').value,
        YearsOfCommercialExperience: document.querySelector('#YearsOfCommercialExperience').value,
        CurrentPosition: document.querySelector('#CurrentPosition').value,
        DesiredPosition: document.querySelector('#DesiredPosition').value,
        stack: document.querySelector('#stack').value
    }
    let err = validateUser(userdata)
    if (err != "") {
        alert(err)
        return
    }

    const intro = document.querySelector('.intro')
    intro.classList.toggle("thin")
    
    const chat = document.querySelector('.chat')
    chat.classList.toggle("thin")

    document.querySelectorAll(".filler").forEach(e => {e.classList.toggle("thin")})

    
    ws.send("start,"+ userdata.username + "," + userdata.YearsOfCommercialExperience + "," + userdata.CurrentPosition + "," + userdata.DesiredPosition + "," + userdata.stack  )
    console.log("starting chat for", userdata);
}

document.querySelector("body").innerHTML+= "<hr>js loaded";
const fillForm = (name) => {
    if (name == 'ilyas') {
        document.querySelector('#username').value = 'Mavlodi'
        document.querySelector('#YearsOfCommercialExperience').value = '1'
        document.querySelector('#CurrentPosition').value = 'Junior'
        document.querySelector('#DesiredPosition').value = 'Strong Junior'
        document.querySelector('#stack').value = 'Golang'
    }
    if (name == 'eldana') {
        document.querySelector('#username').value = 'Karen'
        document.querySelector('#YearsOfCommercialExperience').value = '3'
        document.querySelector('#CurrentPosition').value = 'Middle'
        document.querySelector('#DesiredPosition').value = 'Senior'
        document.querySelector('#stack').value = 'iOS'
    }
}

const validateUser = (userdata) => {
    let errValidateMsg = "Please, enter "
    let arrValidate = []

    if (userdata.username == "") {
        arrValidate.push("name")
    }
    if (userdata.YearsOfCommercialExperience == "") {
        arrValidate.push("years of commercial experience")
    }
    if (userdata.CurrentPosition == "") {
        arrValidate.push("current position")
    }
    if (userdata.DesiredPosition == "") {
        arrValidate.push("desired position")
    }
    if (userdata.stack == "") {
        arrValidate.push("background experience")
    }
    if (arrValidate.length > 0) {
        errValidateMsg += arrValidate.join(", ")
        return errValidateMsg
    }
    return ""
}