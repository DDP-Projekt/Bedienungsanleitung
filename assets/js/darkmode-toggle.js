if (window.localStorage.getItem("dark-mode") === "false") {
	document.querySelector('html').setAttribute("data-theme", "light");
	window.localStorage.setItem("dark-mode", "false");
} else {
	document.querySelector('html').setAttribute("data-theme", "dark");
	window.localStorage.setItem("dark-mode", "true");
}

function toggleDarkMode() {
	if (document.querySelector('html').getAttribute('data-theme') === 'dark') {
		document.querySelector('html').setAttribute("data-theme", "light");
		document.querySelector('#titel').src = 'img/title-dark.svg';
		window.localStorage.setItem("dark-mode", "false");
	}
	else {
		document.querySelector('html').setAttribute("data-theme", "dark");
		document.querySelector('#titel').src = 'img/title.svg';
		window.localStorage.setItem("dark-mode", "true");
	}
}