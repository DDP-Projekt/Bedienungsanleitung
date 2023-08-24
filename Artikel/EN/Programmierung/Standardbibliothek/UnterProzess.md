# Duden/UnterProzess Funktionen
<details>
<summary><h2>Programm_Ausfuehren</h2></summary>
<ul>
<pre>
Grundlegende Funktion um Programme auszuführen.

Parameter:
<ProgrammName>: Ein Pfad zu einer Ausführbaren Datei oder ein Befehl, der sich im PATH befindet
<Argumente>: Die Kommandozeilen Argumente für das Auszuführende Programm (ohne den Programm Namen!)
<StandardEingabe>: Die Eingabe für das Programm. Wird in sein Stdin geschrieben.

Rückgabe:
- Der Exit Code des Programms, oder -1 im Falle eines Fehlers
- Fehler werden in <Fehler> gespeichert
- Stdout wird in <StandardAusgabe> gespeichert
- Stderr wird in <StandardFehlerAusgabe> gespeichert

<StandardAusgabe> und <StandardFehlerAusgabe> dürfen dieselbe
Text Referenz sein. Sollte dies der Fall sein enthalten sie
die Kombinierte Ausgabe von Stdout und Stderr.

Achtung!!! Durch Buffering des Programms sind stdout und stderr
vielleicht nicht in der reihenfolge in der sie ausgegeben wurden.

Sollten <StandardAusgabe>, <StandardFehlerAusgabe> oder <Fehler>
nicht gebraucht werden bleiben sie unverändert.
</pre>
	<li>Parameter: <code>ProgrammName</code>, <code>Argumente</code>, <code>Fehler</code>, <code>StandardEingabe</code>, <code>StandardAusgabe</code>, <code>StandardFehlerAusgabe</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text Liste</code>, <code>Text Referenz</code>, <code>Text</code>, <code>Text Referenz</code>, <code>Text Referenz</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Führe das Programm &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt; und der Eingabe &lt;StandardEingabe&gt; aus, und speichere das Ergebniss in &lt;StandardAusgabe&gt; und &lt;StandardFehlerAusgabe&gt; und mögliche Fehler in &lt;Fehler&gt;&#34;</code></li>
	<li><code>&#34;der Rückgabe Wert vom &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt;, der Eingabe &lt;StandardEingabe&gt;, der Ausgabe &lt;StandardAusgabe&gt; und &lt;StandardFehlerAusgabe&gt; und dem Fehler &lt;Fehler&gt;&#34;</code></li>
	<li><code>&#34;dem Rückgabe Wert vom &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt;, der Eingabe &lt;StandardEingabe&gt;, der Ausgabe &lt;StandardAusgabe&gt; und &lt;StandardFehlerAusgabe&gt; und dem Fehler &lt;Fehler&gt;&#34;</code></li>
	<li><code>&#34;den Rückgabe Wert vom &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt;, der Eingabe &lt;StandardEingabe&gt;, der Ausgabe &lt;StandardAusgabe&gt; und &lt;StandardFehlerAusgabe&gt; und dem Fehler &lt;Fehler&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Programm_Ausfuehren_Fehler</h2></summary>
<ul>
<pre>
Wrapper für Programm_Ausfuehren ohne Stdin, Stdout und Stderr.
</pre>
	<li>Parameter: <code>ProgrammName</code>, <code>Argumente</code>, <code>Fehler</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text Liste</code>, <code>Text Referenz</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Führe das Programm &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt; und einem möglichen Fehler &lt;Fehler&gt; aus&#34;</code></li>
	<li><code>&#34;der Rückgabe Wert von &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt; und einem möglichen Fehler &lt;Fehler&gt;&#34;</code></li>
	<li><code>&#34;dem Rückgabe Wert von &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt; und einem möglichen Fehler &lt;Fehler&gt;&#34;</code></li>
	<li><code>&#34;den Rückgabe Wert von &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt; und einem möglichen Fehler &lt;Fehler&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text out ist "".
Gib
Führe das Programm ProgrammName mit den Argumenten Argumente
und der Eingabe "" aus, und speichere das Ergebniss in
out und out und mögliche Fehler in Fehler
zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Programm_Ausfuehren_Ausgabe</h2></summary>
<ul>
<pre>
Wrapper für Programm_Ausfuehren, der die Kombinierte Ausgabe (Stdout + Stderr) zurückgibt.
</pre>
	<li>Parameter: <code>ProgrammName</code>, <code>Argumente</code>, <code>Eingabe</code>, <code>Fehler</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text Liste</code>, <code>Text</code>, <code>Text Referenz</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Ausgabe von &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt;, der Eingabe &lt;Eingabe&gt; und dem möglichen Fehler &lt;Fehler&gt;&#34;</code></li>
	<li><code>&#34;der Ausgabe von &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt;, der Eingabe &lt;Eingabe&gt; und dem möglichen Fehler &lt;Fehler&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text out ist "".
Führe das Programm ProgrammName mit den Argumenten Argumente
und der Eingabe Eingabe aus, und speichere das Ergebniss in
out und out und mögliche Fehler in Fehler.
Gib out zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Programm_Ausfuehren_Einfach</h2></summary>
<ul>
	<li>Parameter: <code>ProgrammName</code>, <code>Argumente</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text Liste</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Führe das Programm &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt; aus&#34;</code></li>
	<li><code>&#34;der Rückgabe Wert von &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt;&#34;</code></li>
	<li><code>&#34;den Rückgabe Wert von &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt;&#34;</code></li>
	<li><code>&#34;dem Rückgabe Wert von &lt;ProgrammName&gt; mit den Argumenten &lt;Argumente&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text out ist "".
Der Text fehler ist "".
Gib den Rückgabe Wert vom 
ProgrammName mit den Argumenten Argumente,
der Eingabe "", 
der Ausgabe out und out 
und dem Fehler fehler zurück.

</code>
</pre>
</details>


