'use strict'

customElements.define('to-do', TODOElement);
const sidebar = document.getElementById("seitenleiste");
const burger = document.getElementById("hamburger");

// automatically close nav on mobile
closeNav();

function closeNav() {
	// only close nav if on mobile
	if (visualViewport.width > 830) {
		return;
	}

	// hide sidebar
	sidebar.classList.toggle('opened', false);
	burger.classList.toggle('opened', false);
}

function toggleNav() {
	sidebar.classList.toggle('opened');
	burger.classList.toggle('opened');
}

// -1 = prev, 1 = next
function goToSurroundingLink(dir) {
	const anchors = document.querySelectorAll('#artikel-links a');
	let currentIndex = -1;

	// Find the index of the current position
	anchors.forEach(function(anchor, index) {
		if (anchor.href === window.location.href) {
			currentIndex = index;
		}
	});

	if (dir === -1) {
		if (currentIndex <= 0) {
			currentIndex = anchors.length;
		}
		window.location.href = anchors[currentIndex - 1].getAttribute('href');
	}
	else if (dir === 1) {
		if (currentIndex >= anchors.length - 1) {
			currentIndex = -1;
		}
		window.location.href = anchors[currentIndex + 1].getAttribute('href');
	}
}
