// Déclaration de mes varibles select depuis mon html,css
let images = document.querySelectorAll('img');
let nbSlide = images.length;
let suivant = document.querySelector('.right');
let precedent = document.querySelector('.left');
let compte = 0;
let play = document.getElementById("btn-play");


function slideSuivante(){
    images[compte].classList.remove('active');

    if(compte < nbSlide - 1){
        compte++;
    } else {
        compte = 0;
    }

    images[compte].classList.add('active');
    
}
suivant.addEventListener('click', slideSuivante);


function slidePrecedente(){
    images[compte].classList.remove('active');

    if(compte > 0){
        compte--;
    } else {
        compte = nbSlide - 1;
    }

    images[compte].classList.add('active');
    
}
precedent.addEventListener('click', slidePrecedente);


function Start(){



    
}

play.addEventListener('click', Start);

play.onclick = Start;