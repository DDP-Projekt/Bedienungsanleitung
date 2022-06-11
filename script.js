document.addEventListener("DOMContentLoaded", () => {
    const md = document.getElementById("md");
    md.setAttribute("src", !window.location.search ? "Artikel/Startseite.md" : "Artikel/" + window.location.search.slice(3) + ".md");

    md.addEventListener('zero-md-ready', () => {
        applySyntaxHighlighting();
    });
});

let sidebarHidden = false;

function openNav() {
    sidebarHidden = false;
    document.getElementById("seitenleiste").style.display = "flex";
    
    const burger = document.getElementById("hamburger");
    burger.style.left = "calc(var(--sidebar-size) - 20px)";
}

function closeNav() {
    sidebarHidden = true;
    document.getElementById("seitenleiste").style.display = "none";
    
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
    Prism.languages['ddp'] = {
        'placeholder': /<[^<>]+>/,
        'comment': /\[[\s\S]*\]/,
        'builtins': /Schreibe (den|die) (Zahl(en)?|Kommazahl(en)?|Buchstaben?|Texte?)|auf eine Zeile/,
        'string-literal': /\".*\"/,
        'char-literal': /\'.*\'/,
        'boolean-literal': /wahr|falsch/,
        'number': /-?\d+(,\d)?/,
        'control-statement': /[Ww]enn|aber|dann|Sonst|[Ss]olange|[Mm]ache|Für|jede|mit Schrittgröße|[Gg]ibt?|zurück|und macht|einen?|mit den Parametern|vom|und kann so benutzt werden/,
        'punctuation': /\.|\:/,
        'operator': /\b(ist|sind|oder|und|nicht|plus|minus|mal|durch|modulo|hoch|. wurzel von|logisch|kontra|gleich|ungleich|kleiner als|größer als|kleiner als, oder|größer als, oder|natürlicher Logarithmus von|Betrag|Stück|von|bis|als|an der Stelle)\b/,
        'type': /nichts|Typ|Funktion|Zahl(en)?|Kommazahl(en)?|Booleans?|Texte?|Buchstaben?|[Dd]er|[Dd]ie|[Dd]as/,
    };
}