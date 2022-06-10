let sidebarHidden = false;

function openNav() {
    sidebarHidden = false;
    document.getElementById("seitenleiste").style.display = "flex";
}

function closeNav() {
    sidebarHidden = true;
    document.getElementById("seitenleiste").style.display = "none";
}

function toggleNav() {
    const pfeil = document.getElementById("pfeil");
    if (sidebarHidden) {
        openNav();
        pfeil.style.left = "calc(var(--sidebar-size) - 20px)";
        pfeil.innerHTML = "←"
    }
    else {
        closeNav();
        pfeil.style.left = 0;
        pfeil.innerHTML = "→"
    }
}