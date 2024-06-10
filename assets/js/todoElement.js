class TODOElement extends HTMLElement {
	constructor() {
		// Always call super first in constructor
		super();

		this.attachShadow({mode: 'open'});

		const div = document.createElement('div');
		const text = document.createElement('p');
		text.innerHTML = "An diesem Artikel muss noch gearbeitet werden!";
		div.appendChild(text);

		this.shadowRoot.append(div);
	}
}

customElements.define('to-do', TODOElement);