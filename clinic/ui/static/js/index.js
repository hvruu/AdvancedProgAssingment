const audio = new Audio("https://www.fesliyanstudios.com/play-mp3/387");
const buttonsa = document.querySelectorAll("button");

buttonsa.forEach(button => {
    button.addEventListener("click", () => {
        audio.play();

    });
});

function submitPurchase(modalId) {
    const formData = {};
    const formElements = document.querySelector(`#${modalId} .purchase-form`).elements;
    const productNumber = document.querySelector(`#${modalId} .purchase-form`).getAttribute('data-product-number');

    for (let i = 0; i < formElements.length; i++) {
        const element = formElements[i];
        if (element.tagName === "INPUT") {
            formData[element.id] = element.value;
        }
    }

    if (!formData.name || !formData.surname || !formData.cardNumber) {
        alert("Please fill in all the fields.");
        return;
    }

    alert(`Successfully purchased product ${productNumber}!`);
}


var buttons = document.querySelectorAll('.btn');
buttons.forEach(function (button) {
    button.addEventListener('mousedown', function () {
        button.style.transform = 'scale(1.01)';
    });

    button.addEventListener('mouseup', function () {
        button.style.transform = 'scale(1)';
    });

    button.addEventListener('mouseenter', function () {
        button.style.borderColor = '#34db66';
        button.style.color = '#34db66';
    });

    button.addEventListener('mouseleave', function () {
        button.style.borderColor = '#42455a';
        button.style.color = '#e0ffff';
        button.style.boxShadow = 'none';
        button.style.transitionDuration = '50ms';
    });

    button.addEventListener('mouseenter', function () {
        button.style.boxShadow = '0 0 10px rgba(91, 255, 129, 0.5);';
        button.style.transitionDuration = '50ms';
    });

    button.addEventListener('mouseleave', function () {
        button.style.boxShadow = 'none';
        button.style.transitionDuration = '350ms';
    });
});
document.addEventListener('keypress', function (event) {
    if (event.key === 'ArrowLeft') {
        document.getElementById('carouselExampleIndicators').carousel('prev');
    } else if (event.key === 'ArrowRight') {
        document.getElementById('carouselExampleIndicators').carousel('next');
    }
});


var cardImages = document.querySelectorAll(".card-img-top");

cardImages.forEach(function (image) {
    image.addEventListener("mouseover", function () {
        image.style.transform = "scale(1.01)";
        image.style.transition = "transform 0.7s ease";
    });
    image.addEventListener("mouseout", function () {
        image.style.transform = "scale(1)";
    });
});



function onEntry(entry) {
    entry.forEach(change => {
        if (change.isIntersecting) {
            change.target.classList.add('element-show');
        }
    });
}
let options = { threshold: [0.5] };
let observer = new IntersectionObserver(onEntry, options);
let elements = document.querySelectorAll('.element-animation');
for (let elm of elements) {
    observer.observe(elm);
}
