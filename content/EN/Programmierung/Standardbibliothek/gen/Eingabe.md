+++
title = "Eingabe"
weight = 1
+++
# Duden/Eingabe functions
<details>
<summary><h2>War_EOF</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn ein EOF Zeichen gelesen wurde.
</pre>
</li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
	<li>Return type: <code>Text</code></li>
</ul>

<h3>Aliases</h3>
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

<details>
<summary><h2>Lies_Token_Gepuffert</h2></summary>
<ul>
<pre>
Wenn Eingabe_Puffer eine Eingabe enthält
wird diese zurückgegeben und Eingabe_Puffer geleert.
Ansonsten wird die nächste Eingabe direkt zurückgegeben.
</pre>
</li>
	<li>Return type: <code>Text</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;der nächste Text&#34;</code></li>
	<li><code>&#34;den nächsten Text&#34;</code></li>
	<li><code>&#34;der naechste Text&#34;</code></li>
	<li><code>&#34;den naechsten Text&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von Eingabe_Puffer größer als 0 ist, dann:
	Der Text Rückgabe ist Eingabe_Puffer.
	Eingabe_Puffer ist "".
	Gib Rückgabe zurück.

Wenn die Benutzereingabe nicht vorbei ist, gib die nächste Eingabe zurück.
Gib "" zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Lies_Zahl_Gepuffert</h2></summary>
<ul>
<pre>
Lies_Token_Gepuffert als Zahl.

Vorher sollte Hat_Zahl überprüft werden.
</pre>
</li>
	<li>Return type: <code>Zahl</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;die nächste Zahl&#34;</code></li>
	<li><code>&#34;die naechste Zahl&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib den nächsten Text als Zahl zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Lies_Kommazahl_Gepuffert</h2></summary>
<ul>
<pre>
Lies_Token_Gepuffert als Kommazahl.

Vorher sollte Hat_Zahl überprüft werden.
</pre>
</li>
	<li>Return type: <code>Kommazahl</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;die nächste Kommazahl&#34;</code></li>
	<li><code>&#34;die naechste Kommazahl&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib den nächsten Text als Kommazahl zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Lies_Buchstabe_Gepuffert</h2></summary>
<ul>
<pre>
Wenn Eingabe_Puffer eine Eingabe enthält wird
der erste Buchstabe davon zurückgegeben und Eingabe_Puffer
entsprechend angepasst.

Wenn Gepufferter_Buchstabe einen Buchstaben enthält wird dieser
zurückgeben und Gepufferter_Buchstabe auf 0 gesetzt.

Ansonsten wird das nächste Zeichen direkt gelesen.
</pre>
</li>
	<li>Return type: <code>Buchstabe</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;der nächste Buchstabe&#34;</code></li>
	<li><code>&#34;der naechste Buchstabe&#34;</code></li>
	<li><code>&#34;den nächsten Buchstaben&#34;</code></li>
	<li><code>&#34;den naechsten Buchstaben&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von Eingabe_Puffer ungleich 0 ist, dann:
	Der Buchstabe Rückgabe ist Eingabe_Puffer an der Stelle 1.
	Speichere Eingabe_Puffer von 2 bis (die Länge von Eingabe_Puffer) in Eingabe_Puffer.
	Gib Rückgabe zurück.

Wenn Gepufferter_Buchstabe als Zahl ungleich 0 ist, dann:
	Der Buchstabe Rückgabe ist Gepufferter_Buchstabe.
	Speichere 0 als Buchstabe in Gepufferter_Buchstabe.
	Gib Rückgabe zurück.

Gib das nächste Zeichen aus der Standardeingabe zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Lies_Zeile_Gepuffert</h2></summary>
<ul>
<pre>
Ignoriert vorrangehende Leerzeichen und gibt
dann die ganze Zeile bis zu, aber ohne '\n' oder EOF zurück.

'\n' bzw. EOF wird dabei nicht nicht in einen Puffer gelegt
sondern verworfen.
</pre>
</li>
	<li>Return type: <code>Text</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;die nächste Zeile&#34;</code></li>
	<li><code>&#34;die naechste Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text zeile ist "".
Wenn die Länge von Eingabe_Puffer ungleich 0 ist, dann:
	Füge Eingabe_Puffer an zeile an.
	Eingabe_Puffer ist "".
	Wenn Gepufferter_Buchstabe als Zahl ungleich 0 ist, Füge Gepufferter_Buchstabe an zeile an.
	Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.
Sonst:
	Solange Gepufferter_Buchstabe ungleich '\n' ist und (Gepufferter_Buchstabe ein Leerzeichen ist oder Gepufferter_Buchstabe ein Kontrollzeichen ist), mache:
		Wenn die Benutzereingabe zu Ende gewesen ist, gib "" zurück.
		Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.

Solange Gepufferter_Buchstabe ungleich '\n' ist und die Benutzereingabe nicht vorbei ist, mache:
	Speichere zeile verkettet mit Gepufferter_Buchstabe in zeile.
	Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.

Speichere 0 als Buchstabe in Gepufferter_Buchstabe.
Gib zeile zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Hat_Token</h2></summary>
<ul>
<pre>
Gibt an ob etwas mit einer der obigen Funktionen
gelesen werden kann.
</pre>
</li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;eine Eingabe vorhanden ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von Eingabe_Puffer größer als 0 ist, gib wahr zurück.
Wenn die Benutzereingabe zu Ende gewesen ist, gib falsch zurück.
Speichere den nächsten Text in Eingabe_Puffer.
Gib [wahr wenn] die Länge von Eingabe_Puffer größer als 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Hat_Zahl</h2></summary>
<ul>
<pre>
Gibt an ob die nächste Eingabe (direkt oder im Puffer)
zu einer Zahl umgewandelt werden kann.
</pre>
</li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;die nächste Eingabe eine Zahl ist&#34;</code></li>
	<li><code>&#34;die naechste Eingabe eine Zahl ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib eine Eingabe vorhanden ist und Eingabe_Puffer eine Zahl ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Hat_Buchstabe</h2></summary>
<ul>
<pre>
Gibt an ob ein Buchstabe (direkt oder aus dem Puffer)
gelesen werden kann.
</pre>
</li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;ein Buchstabe vorhanden ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib nicht war_eof oder die Länge von Eingabe_Puffer ungleich 0 ist oder Gepufferter_Buchstabe als Zahl ungleich 0 ist zurück.

</code>
</pre>
</details>


