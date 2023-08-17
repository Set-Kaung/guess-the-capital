const questionContainer = document.querySelector(".question span")
const enterBtn = document.querySelector("button.enter")
const answerBox = document.querySelector("#answer")
const againBtn = document.querySelector(".again")
const answerArea = document.querySelector(".ans")

class Question{
    constructor(country, capital){
        this.country = country
        this.capital = capital
    }
}
var Ques 

function renderQuestion(){
    if(answerArea.classList.contains("hide")){
        answerArea.classList.remove("hide")
    }
    if(enterBtn.classList.contains("hide")){
        enterBtn.classList.remove("hide")
    }
    questionContainer.textContent = `What is the capital of ${Ques.country}?` 
}

async function getQuestion() {
    const response = await fetch("http://127.0.0.1:4444/question", { mode: 'cors'});
    const question = await response.json();
    Ques = new Question(question.country, question.capital)
    renderQuestion()
}

function getAnswer(){
    let ans = answerBox.value
    answerBox.value = ""
    checkAnswer(ans)
}

function checkAnswer(answer){
    answer = answer.replaceAll(" ", "")
    let country = Ques.country
    let capital = Ques.capital

    if(answer.toLowerCase() == capital.toLowerCase()){
        questionContainer.textContent = `You are right. The capital of ${country} is ${capital}.`
    }else{
        questionContainer.textContent = `Sorry. The capital of ${country} is ${capital}.`
    }
    enterBtn.classList.toggle("hide")
    againBtn.classList.toggle("hide")
    answerArea.classList.toggle("hide")
}

function playAgain(){
    againBtn.classList.toggle("hide")
    getQuestion()
}



enterBtn.addEventListener("click", getAnswer)
againBtn.addEventListener("click", playAgain)
document.addEventListener("keyup", e =>{
    if(e.key == 13){
        getAnswer()
    }
})

window.onload = getQuestion



