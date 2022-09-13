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
	populateListItems();

	populateMDElement();

    const md = document.getElementById("md"); // zero-md element
    // apply syntax highlighting once zero-md and Prism.js are loaded
    md.addEventListener('zero-md-ready', () => {
        applySyntaxHighlighting();
    });

    closeNav(); // automatically close nav
});

function populateListItems() {
	let searchParams = new URLSearchParams(window.location.search);
    let articleLinks = document.getElementById('article-links')

    traverseListElement(articleLinks, function (element) {
        // only allow <a> tags without a href attribute
        if (element.nodeName !== "A" || element.getAttribute("href")) {
            return;
        }

        // add href attribute
        element.setAttribute("href", `?p=${path}${element.innerHTML}&lang=${searchParams.get('lang')}`)
    })
}

let path = "";

/**
 * Super janky lol
 * @param {Element} element The Element
 * @param {function(Element)} callback
 */
function traverseListElement(element, callback) {
    // only allow elements with children
    if (element == null || element.children.length === 0) {
        return;
    }

    // loop through children
    for (const child of element.children) {
        if (child.nodeName === "OL") {
            // add the innerHTML of the <a> tag above the <ol> tag to the path
            path += child.parentElement.firstElementChild.innerHTML + "/";
        }

        callback(child);
        traverseListElement(child, callback);
    }

    // remove old path
    if (element === element.parentElement.lastElementChild) {
        path = path.replace(element.parentElement.firstElementChild.innerHTML + "/", "");
    }
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
    document.getElementById("seitenleiste").style.display = "flex";
    
    // push the hamburger button to the end of the sidebar.
    const burger = document.getElementById("hamburger");
    burger.style.left = "calc(var(--sidebar-size) - 15px)";
}

function closeNav() {
    // only close nav if on mobile
    if (visualViewport.width > 830) {
        return;
    }

    // hide sidebar
    sidebarHidden = true;
    document.getElementById("seitenleiste").style.display = "none";
    
    // set hamburger button to the left of the screen
    const burger = document.getElementById("hamburger");
    burger.style.left = "0";
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
        'control-statement': /[Ww]enn|aber|dann|Sonst|[Ss]olange|[Mm]ache|Für|jede|mit Schrittgröße|[Gg]ibt?|zurück|macht|einen?|mit den Parametern|vom|und kann so benutzt werden|Binde|ein/,
        'punctuation': /[.:]/,
        'operator': /\b(ist|sind|oder|und|nicht|plus|minus|mal|durch|modulo|hoch|. wurzel von|logisch|kontra|gleich|ungleich|kleiner als|größer als|kleiner als, oder|größer als, oder|natürlicher Logarithmus von|Betrag|Stück|von|bis|als|an der Stelle)\b/,
        'type': /nichts|Typ|Funktion|Zahl(en)?|Kommazahl(en)?|Booleans?|Texte?|Buchstaben?|[Dd]er|[Dd]ie|[Dd]as/,
    };
}

function changeLanguage(lang) {
	const languages = ['DE', 'EN']

    const searchParams = new URLSearchParams(window.location.search);
    searchParams.set("lang", languages[lang]);
    window.location.search = searchParams.toString();
}
