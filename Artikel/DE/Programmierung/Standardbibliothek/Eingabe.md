# Duden/Eingabe Funktionen
<details>
<summary><h2>War_EOF</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn ein EOF Zeichen gelesen wurde.
</pre>
</li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Benutzereingabe zu Ende gewesen ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib war_eof zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Ist_Nicht_EOF</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn noch kein EOF Zeichen gelesen wurde.
</pre>
</li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Benutzereingabe nicht vorbei ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib nicht war_eof zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Benutzereingabe_Zuruecksetzten</h2></summary>
<ul>
<pre>
Setzt die interne war_eof variable auf falsch.
</pre>
</li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;setzte die Benutzereingabe zurück&#34;</code></li>
	<li><code>&#34;Setzte die Benutzereingabe zurück&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
war_eof ist falsch.

</code>
</pre>
</details>

<details>
<summary><h2>Lies_Token_in_Puffer</h2></summary>
<ul>
<pre>
Liest eine einzelne Eingabe in Eingabe_Puffer.

Dabei werden vorrangehende Leerzeichen ignoriert
und ein einzelnes darrauffolgendes Leerzeichen in
Gepufferter_Buchstabe gespeichert
</pre>
</li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die nächste Eingabe&#34;</code></li>
	<li><code>&#34;die naechste Eingabe&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
[vorangehende Leerstellen überspringen]
Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.
Solange Gepufferter_Buchstabe ein Leerzeichen ist oder Gepufferter_Buchstabe ein Kontrollzeichen ist, mache:
	Wenn die Benutzereingabe zu Ende gewesen ist, gib "" zurück.
	Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.

[Text bis Leerstelle Lesen]
Der Text token ist "".
Solange nicht (Gepufferter_Buchstabe ein Leerzeichen ist oder Gepufferter_Buchstabe ein Kontrollzeichen ist), mache:
	Speichere token verkettet mit Gepufferter_Buchstabe in token.
	Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.

Gib token zurück.

</code>
</pre>
</details>


