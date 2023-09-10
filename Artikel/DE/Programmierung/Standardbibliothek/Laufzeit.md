# Duden/Laufzeit Funktionen
<details>
<summary><h2>Programm_Beenden</h2></summary>
<ul>
<pre>
Beendet das Programm mit dem gegebenen Code.
</pre>
	<li>Parameter: <code>Code</code></li>
	<li>Parameter Typ: <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Beende das Programm mit Code &lt;Code&gt;&#34;</code></li>
	<li><code>&#34;beende das Programm mit Code &lt;Code&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Programm_Beenden_Standard</h2></summary>
<ul>
<pre>
Beendet das Programm mit Code 0.
</pre>
</li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Beende das Programm&#34;</code></li>
	<li><code>&#34;beende das Programm&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Beende das Programm mit Code 0.

</code>
</pre>
</details>

<details>
<summary><h2>Laufzeitfehler</h2></summary>
<ul>
<pre>
Wirft einen Laufzeitfehler mit einer Nachricht und einem Code.
</pre>
	<li>Parameter: <code>Nachricht</code>, <code>Code</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Löse einen Laufzeitfehler mit der Nachricht &lt;Nachricht&gt; und dem Code &lt;Code&gt; aus&#34;</code></li>
	<li><code>&#34;löse einen Laufzeitfehler mit der Nachricht &lt;Nachricht&gt; und dem Code &lt;Code&gt; aus&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Befehlszeilenargumente</h2></summary>
<ul>
<pre>
Gibt eine Text Liste zurück die alle übergebenen Befehlszeilenargumente enthält.
Index 1 enthält immer den Programmpfad.
</pre>
</li>
	<li>Rückgabe Typ: <code>Text Liste</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Befehlszeilenargumente&#34;</code></li>
	<li><code>&#34;den Befehlszeilenargumenten&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpruntime.a"</code>
</details>

<details>
<summary><h2>Betriebssystem</h2></summary>
<ul>
<pre>
Gibt je nach Betriebssystem entweder "Windows" oder "Linux" zurück.
</pre>
</li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;das Betriebssystem&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpruntime.a"</code>
</details>

<details>
<summary><h2>Ist_Befehlszeile</h2></summary>
<ul>
<pre>
Entspricht C's `isatty` funktion.
</pre>
</li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Benutzereingabe eine Befehlszeile ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Macht_Nichts</h2></summary>
<ul>
<pre>
Leere Funktion, die nichts macht und nichts zurückgibt.
Nützlich in einigen wenigen Fällen.
</pre>
</li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Tue nichts&#34;</code></li>
	<li><code>&#34;tue nichts&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">

</code>
</pre>
</details>


