# Duden/Ausgabe Funktionen
<details>
<summary><h2>Schreibe_Zahlen_Liste_Getrennt</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Zahlen_Liste_Getrennt schreibt alle Elemente einer gegebenen Zahlen Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>liste</code>, <code>seperator</code></li>
	<li>Parameter Typen: <code>Zahlen Liste</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Zahlen Liste &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste größer als 0 ist, dann:
	Wenn die Länge von liste größer als 1 ist, dann:
		Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
			Schreibe (liste an der Stelle i).
			Schreibe seperator.
	Schreibe (liste an der Stelle (die Länge von liste)).

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Kommazahlen_Liste_Getrennt</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Kommazahlen_Liste_Getrennt schreibt alle Elemente einer gegebenen Kommazahlen Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>liste</code>, <code>seperator</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Kommazahlen Liste &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste größer als 0 ist, dann:
	Wenn die Länge von liste größer als 1 ist, dann:
		Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
			Schreibe (liste an der Stelle i).
			Schreibe seperator.
	Schreibe (liste an der Stelle (die Länge von liste)).

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Boolean_Liste_Getrennt</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Boolean_Liste_Getrennt schreibt alle Elemente einer gegebenen Boolean Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>liste</code>, <code>seperator</code></li>
	<li>Parameter Typen: <code>Boolean Liste</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Boolean Liste &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste größer als 0 ist, dann:
	Wenn die Länge von liste größer als 1 ist, dann:
		Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
			Schreibe (liste an der Stelle i).
			Schreibe seperator.
	Schreibe (liste an der Stelle (die Länge von liste)).

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Buchstaben_Liste</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Buchstaben_Liste schreibt alle Elemente einer Buchstaben Liste getrennt mit einem Komma in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Buchstaben Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Buchstaben Liste &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Boolean_Liste_Zeile</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Boolean_Liste schreibt alle Elemente einer Boolean Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Boolean Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Boolean Liste &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Text_Liste_Zeile</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Text_Liste schreibt alle Elemente einer Text Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Text Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Text Liste &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>SchreibeZeile_Boolean</h2></summary>
<ul>
<pre>
Die Funktion SchreibeZeile_Boolean schreibt einen gegebenen Boolean (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Boolean</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe den Boolean &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1.
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Kommazahl</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Kommazahl schreibt eine gegebene Kommazahl (p1) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Kommazahl &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Schreibe_Boolean</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Boolean schreibt einen gegebenen Boolean (p1) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Boolean</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe den Boolean &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Schreibe_Fehler</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Text schreibt einen gegebenen Text (fehler) in den Standart Error Stream.
</pre>
	<li>Parameter: <code>fehler</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe den Fehler &lt;fehler&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>SchreibeZeile_Zahl</h2></summary>
<ul>
<pre>
Die Funktion SchreibeZeile_Zahl schreibt eine gegebene Zahl (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Zahl &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1.
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>SchreibeZeile_Buchstabe</h2></summary>
<ul>
<pre>
Die Funktion SchreibeZeile_Buchstabe schreibt einen gegebenen Buchstaben (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe den Buchstaben &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1.
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Text_Liste_Getrennt</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Text_Liste_Getrennt schreibt alle Elemente einer gegebenen Text Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>liste</code>, <code>seperator</code></li>
	<li>Parameter Typen: <code>Text Liste</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Text Liste &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste größer als 0 ist, dann:
	Wenn die Länge von liste größer als 1 ist, dann:
		Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
			Schreibe (liste an der Stelle i).
			Schreibe seperator.
	Schreibe (liste an der Stelle (die Länge von liste)).

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Zahlen_Liste</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Zahlen_Liste schreibt alle Elemente einer Zahlen Liste getrennt mit einem Komma in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Zahlen Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Zahlen Liste &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Zahl</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Zahl schreibt eine gegebene Zahl (p1) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Zahl &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Schreibe_Kommazahlen_Liste</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Kommazahlen_Liste schreibt alle Elemente einer Kommazahlen Liste getrennt mit einem Komma in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Kommazahlen Liste &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Boolean_Liste</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Boolean_Liste schreibt alle Elemente einer Boolean Liste getrennt mit einem Komma in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Boolean Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Boolean Liste &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Zahlen_Liste_Zeile</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Zahlen_Liste schreibt alle Elemente einer Zahlen Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Zahlen Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Zahlen Liste &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Kommazahlen_Liste_Zeile</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Kommazahlen_Liste schreibt alle Elemente einer Kommazahlen Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Kommazahlen Liste &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Buchstaben_Liste_Getrennt</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Buchstaben_Liste_Getrennt schreibt alle Elemente einer gegebenen Buchstaben Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>liste</code>, <code>seperator</code></li>
	<li>Parameter Typen: <code>Buchstaben Liste</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Buchstaben Liste &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;liste&gt; mit dem Seperator &lt;seperator&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste größer als 0 ist, dann:
	Wenn die Länge von liste größer als 1 ist, dann:
		Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
			Schreibe (liste an der Stelle i).
			Schreibe seperator.
	Schreibe (liste an der Stelle (die Länge von liste)).

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Text</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Text schreibt einen gegebenen Text (p1) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe den Text &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>SchreibeZeile_Kommazahl</h2></summary>
<ul>
<pre>
Die Funktion SchreibeZeile_Kommazahl schreibt eine gegebene Kommazahl (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Kommazahl &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1.
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>SchreibeZeile_Text</h2></summary>
<ul>
<pre>
Die Funktion SchreibeZeile_Text schreibt einen gegebenen Text (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe den Text &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1.
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Text_Liste</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Text_Liste schreibt alle Elemente einer Text Liste getrennt mit einem Komma in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Text Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Text Liste &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Buchstaben_Liste_Zeile</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Buchstaben_Liste schreibt alle Elemente einer Buchstaben Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Buchstaben Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe die Buchstaben Liste &lt;p1&gt; auf eine Zeile&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt; auf eine Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe p1 mit dem Seperator ", ".
Schreibe '\n'.

</code>
</pre>
</details>

<details>
<summary><h2>Schreibe_Buchstabe</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Buchstabe schreibt einen gegebenen Buchstaben (p1) in den Standart Output Stream.
</pre>
	<li>Parameter: <code>p1</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe den Buchstaben &lt;p1&gt;&#34;</code></li>
	<li><code>&#34;Schreibe &lt;p1&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>


