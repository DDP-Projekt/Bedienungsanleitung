document.addEventListener("DOMContentLoaded", () => {
    // get all a tags inside an ordered list item.
    const links = document.querySelectorAll("ol li a");
    // if there is no href tag defined, add one using the contents of the a tag.
    links.forEach(x => !x.getAttribute("href") ? x.setAttribute("href", "?p=" + x.innerHTML.replace("&amp;", "&")) : x);

    const md = document.getElementById("md"); // zero-md element
    // point source to the correct md file using the query contents. If the query is empty, use the start page.
    md.setAttribute("src", !window.location.search ? "Artikel/Startseite.md" : "Artikel/" + window.location.search.slice(3) + ".md");

    // apply syntax highlighting once zero-md and Prism.js are loaded
    md.addEventListener('zero-md-ready', () => {
        applySyntaxHighlighting();
    });
});

let sidebarHidden = false;

function openNav() {
    // show sidebar
    sidebarHidden = false;
    document.getElementById("seitenleiste").style.display = "flex";
    
    // push the hamburger button to the end of the sidebar.
    const burger = document.getElementById("hamburger");
    burger.style.left = "calc(var(--sidebar-size) - 20px)";
}

function closeNav() {
    // hide sidebar
    sidebarHidden = true;
    document.getElementById("seitenleiste").style.display = "none";
    
    // set hamburger button to the left of the screen
    const burger = document.getElementById("hamburger");
    burger.style.left = "10px";
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
        'comment': /\[[\s\S]*?\]/,
        'builtins': /Schreibe (den|die) (Zahl(en)?|Kommazahl(en)?|Buchstaben?|Texte?)|auf eine Zeile/,
        'string-literal': /\".*\"/,
        'char-literal': /\'.*\'/,
        'boolean-literal': /wahr|falsch/,
        'number': /-?\d+(,\d)?/,
        'control-statement': /[Ww]enn|aber|dann|Sonst|[Ss]olange|[Mm]ache|Für|jede|mit Schrittgröße|[Gg]ibt?|zurück|macht|einen?|mit den Parametern|vom|und kann so benutzt werden/,
        'punctuation': /\.|\:/,
        'operator': /\b(ist|sind|oder|und|nicht|plus|minus|mal|durch|modulo|hoch|. wurzel von|logisch|kontra|gleich|ungleich|kleiner als|größer als|kleiner als, oder|größer als, oder|natürlicher Logarithmus von|Betrag|Stück|von|bis|als|an der Stelle)\b/,
        'type': /nichts|Typ|Funktion|Zahl(en)?|Kommazahl(en)?|Booleans?|Texte?|Buchstaben?|[Dd]er|[Dd]ie|[Dd]as/,
    };
}