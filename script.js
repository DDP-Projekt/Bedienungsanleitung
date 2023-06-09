'use strict'

const searchParams = new URLSearchParams(window.location.search);
if (searchParams.get('p') === null) {
	searchParams.set('p', 'Startseite')
	window.location.search = searchParams.toString();
}

if (searchParams.get('lang') === null) {
	searchParams.set('lang', 'DE')
	window.location.search = searchParams.toString();
}

customElements.define('to-do', TODOElement);
document.addEventListener("DOMContentLoaded", () => {
	const details = document.querySelectorAll('#artikel-links details');
	for (const detail of details) {
		detail.addEventListener('toggle', (event) => {
			window.localStorage.setItem(event.target.id, event.target.open)
		})

		if (window.localStorage.getItem(detail.id) === "true") {
			detail.open = true
		}
	}

	if (window.localStorage.getItem("dark-mode") === "true") {
		document.querySelector('#dark').media = "all"
		document.querySelector('#light').media = "not all"
	} else {
		document.querySelector('#dark').media = "not all"
		document.querySelector('#light').media = "all"
	}

	closeNav(); // automatically close nav

	changeLangSelectFlag();

	populateListItems()

	populateMDElement();

	const md = document.getElementById("md"); // zero-md element
	// apply syntax highlighting once zero-md and Prism.js are loaded
	md.addEventListener('zero-md-ready', () => {
		applySyntaxHighlighting();
	});
});

function populateListItems() {
	const anchors = document.querySelectorAll('#artikel-links a');

	anchors.forEach(function(anchor) {
		const parentPath = getParentPath(anchor);
		const content = anchor.textContent.trim();
		let href = '';

		if (parentPath !== "Die Deutsche Programmiersprache") {
			href = `?p=${parentPath}/${content}`;
		} else {
			href = `?p=${content}`;
		}

		anchor.setAttribute('href', href + `&lang=${searchParams.get('lang')}`);
	});
}

function getParentPath(element) {
	if (!element || !element.parentNode) {
		return '';
	}

	let parent = element.parentNode;
	let parentPath = '';

	while (parent) {
		if (parent.nodeName === 'OL') {
			const parentAnchor = parent.previousElementSibling.querySelector('a');
			const parentContent = parentAnchor ? parentAnchor.textContent.trim() : '';
			const grandParentPath = getParentPath(parentAnchor);
			parentPath = grandParentPath ? `${grandParentPath}/${parentContent}` : parentContent;
			break;
		}

		parent = parent.parentNode;
	}

	return parentPath;
}

function populateMDElement() {
	const searchParams = new URLSearchParams(window.location.search);
	const destinationPage = "Artikel/" + searchParams.get('lang') + '/' + searchParams.get('p') + ".md";

	const md = document.getElementById("md"); // zero-md element
	// point source to the correct md file using the query contents. If the query is empty, use the start page.
	md.setAttribute("src", searchParams.get('p') === null ? "Startseite" : destinationPage);
}

let sidebarHidden = false;

function openNav() {
	// show sidebar
	sidebarHidden = false;

	const sidebar = document.getElementById("seitenleiste");
	sidebar.classList.toggle('opened', true)

	const burger = document.getElementById("hamburger");
	burger.classList.toggle('opened', true)
}

function closeNav() {
	// only close nav if on mobile
	if (visualViewport.width > 830) {
		return;
	}

	// hide sidebar
	sidebarHidden = true;

	const sidebar = document.getElementById("seitenleiste");
	sidebar.classList.toggle('opened', false)

	const burger = document.getElementById("hamburger");
	burger.classList.toggle('opened', false)
}

function toggleNav() {
	if (sidebarHidden) {
		openNav();
	}
	else {
		closeNav();
	}
}

function applySyntaxHighlighting() {
	// regex rules for syntax highlighting
	Prism.languages['ddp'] = {
		'placeholder': /<[^<>]+>/,
		'comment': /\[[\s\S]*?]/,
		'builtins': /Schreibe (den|die) (Zahl(en)?|Kommazahl(en)?|Buchstaben?|Texte?)|auf eine Zeile/,
		'string-literal': /".*"/,
		'char-literal': /'.*'/,
		'boolean-literal': /wahr|falsch/,
		'number': /-?\d+(,\d)?/,
		'control-statement': /[Ww]enn|aber|dann|Sonst|[Ss]olange|[Mm]ache|[Ww]iederhole|Für|jede|mit Schrittgröße|[Gg]ibt?|zurück|macht|einen?|mit den Parametern|vom|und kann so benutzt werden|Binde|ein/,
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

function changeLangSelectFlag() {
	const searchParams = new URLSearchParams(window.location.search);
	const lang = searchParams.get("lang");

	const elm = document.getElementById('sprach-toggle');
	if (lang === "DE") {
		elm.setAttribute('src', 'img/flags/EN.png')
	}
	else {
		elm.setAttribute('src', 'img/flags/DE.png')
	}
}

function toggleLang() {
	const searchParams = new URLSearchParams(window.location.search);
	const lang = searchParams.get("lang");

	if (lang === "DE") {
		searchParams.set("lang", "EN");
	}
	else {
		searchParams.set("lang", "DE");
	}
	window.location.search = searchParams.toString();
}

function toggleDarkMode() {
	let isDark = document.querySelector('#dark').media === 'all'
	if (isDark) {
		document.querySelector('#dark').media = 'not all'
		document.querySelector('#light').media = 'all'
		window.localStorage.setItem("dark-mode", "false")
	}
	else {
		document.querySelector('#dark').media = 'all'
		document.querySelector('#light').media = 'not all'
		window.localStorage.setItem("dark-mode", "true")
	}
}

// -1 = prev, 1 = next
function goToSurroundingLink(dir) {
	const anchors = document.querySelectorAll('#artikel-links a');
	let currentIndex = -1;

	// Find the index of the current position
	anchors.forEach(function(anchor, index) {
		if (anchor.getAttribute('href') === decodeURIComponent(window.location.search)) {
			currentIndex = index;
		}
	});

	if (dir === -1) {
		if (currentIndex <= 0) {
			currentIndex = anchors.length
		}
		window.location.href = anchors[currentIndex - 1].getAttribute('href')
	}
	else if (dir === 1) {
		if (currentIndex >= anchors.length - 1) {
			currentIndex = -1
		}
		window.location.href = anchors[currentIndex + 1].getAttribute('href')
	}
}
