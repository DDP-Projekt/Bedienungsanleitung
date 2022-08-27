class TODOElement extends HTMLElement {
	constructor() {
		// Always call super first in constructor
		super();


		this.attachShadow({mode: 'open'});

		const div = document.createElement('div');
		div.style = `
						width: 99%;
						display: flex;
						height: 50px;
						border: 5px solid #df5e5e;
						background-color: #303336;
						justify-content: center;
						align-content: center;
						border-radius: 15px;
					`;



		const text = document.createElement('p');
		text.innerHTML = "An diesem Artikel muss noch gearbeitet werden!";
		div.appendChild(text);

		this.shadowRoot.append(div);
	}
}