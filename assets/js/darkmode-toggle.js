if (window.localStorage.getItem("dark-mode") === "false") {
	document.querySelector('#dark').media = "not all";
	document.querySelector('#light').media = "all";
} else {
	document.querySelector('#dark').media = "all";
	document.querySelector('#light').media = "not all";
}

function toggleDarkMode() {
	let isDark = document.querySelector('#dark').media === 'all';
	if (isDark) {
		document.querySelector('#dark').media = 'not all';
		document.querySelector('#light').media = 'all';
		window.localStorage.setItem("dark-mode", "false");
	}
	else {
		document.querySelector('#dark').media = 'all';
		document.querySelector('#light').media = 'not all';
		window.localStorage.setItem("dark-mode", "true");
	}
}