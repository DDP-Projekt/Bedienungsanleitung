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
