'use strict'

customElements.define('to-do', TODOElement);
document.addEventListener("DOMContentLoaded", () => {
	closeNav(); // automatically close nav
});

let sidebarHidden = false;

function openNav() {
	// show sidebar
	sidebarHidden = false;

	const sidebar = document.getElementById("seitenleiste");
	sidebar.classList.toggle('opened', true);

	const burger = document.getElementById("hamburger");
	burger.classList.toggle('opened', true);
}

function closeNav() {
	// only close nav if on mobile
	if (visualViewport.width > 830) {
		return;
	}

	// hide sidebar
	sidebarHidden = true;

	const sidebar = document.getElementById("seitenleiste");
	sidebar.classList.toggle('opened', false);

	const burger = document.getElementById("hamburger");
	burger.classList.toggle('opened', false);
}

function toggleNav() {
	if (sidebarHidden) {
		openNav();
	}
	else {
		closeNav();
	}
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
