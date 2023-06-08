# Duden/Datei_Eingabe Funktionen
<details>
<summary><h2>Lies_Text_Datei</h2></summary>
<ul>
<pre>
Die Funktion Lies_Text_Datei speichert den Inhalt der Datei, die an dem gegebenen Pfad liegt, in ref und gibt die Anzahl der Bytes der gelesenen Datei zurück.
Wenn ein Fehler auftreten sollte, ist der zurückgegebene Wert negativ und die Fehler meldung in ref geschrieben
</pre>
	<li>Parameter: <code>Pfad</code>, <code>ref</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text Referenz</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lies den Text in &lt;Pfad&gt; und speichere ihn in &lt;ref&gt;&#34;</code></li>
	<li><code>&#34;die Anzahl der Bytes, die aus &lt;Pfad&gt; gelesen und in &lt;ref&gt; gespeichert wurden&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>


