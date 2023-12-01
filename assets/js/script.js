'use strict'

const searchParams = new URLSearchParams(window.location.search);

customElements.define('to-do', TODOElement);
document.addEventListener("DOMContentLoaded", () => {
	closeNav(); // automatically close nav

	//const md = document.getElementById("md"); // zero-md element
	// apply syntax highlighting once zero-md and Prism.js are loaded
	/*md.addEventListener('zero-md-ready', () => {
		applySyntaxHighlighting();
	});*/
});

function applySyntaxHighlighting() {
	// regex rules for syntax highlighting
	Prism.languages['ddp'] = {
		'placeholder': /<[^<>]+>/,
		'comment': /\[[\s\S]*?]/,
		'builtins': /Schreibe\s*(den|die)?\s*(Zahl(en)?|Kommazahl(en)?|Buchstaben?|Texte?)?|auf eine Zeile/,
		'string-literal': /".*"/,
		'char-literal': /'.*'/,
		'boolean-literal': /wahr|falsch/,
		'number': /-?\d+(,\d)?/,
		'control-statement': /[Ww]enn|aber|dann|Sonst|[Ss]olange|[Mm]ache|[Ww]iederhole|Für|jede|mit Schrittgröße|[Gg]ibt?|zurück|macht|einen?|mit den Parametern|vom|[Uu]nd kann so benutzt werden|Binde|ein/,
		'punctuation': /[.:]/,
		'operator': /\b(ist|sind|oder|und|nicht|plus|minus|mal|durch|modulo|hoch|. wurzel von|logisch|kontra|gleich|ungleich|kleiner als|größer als|kleiner als, oder|größer als, oder|natürlicher Logarithmus von|Betrag|Stück|von|bis|als|an der Stelle)\b/,
		'type': /nichts|Typ|Funktion|Zahl(en)?|Kommazahl(en)?|Booleans?|Texte?|Buchstaben?|[Dd]er|[Dd]ie|[Dd]as/,
	};

	Prism.languages['terminal'] = {
		'placeholder': /<[^<>]+>/,
		'file': /(\w+)\.(\w+)/,
		'command': /(?<=\$\s+)(\w+)|\.\/(\w+)/,
		'option': /(\w+)/,
		'string-literal': /".*"/,
	};
}

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
