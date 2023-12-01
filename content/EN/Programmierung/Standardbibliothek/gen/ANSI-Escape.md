+++
title = "ANSI-Escape"
weight = 1
+++
# Duden/ANSI-Escape functions
<details>
<summary><h2>ANSI_Funktion_Ausführen_1Arg</h2></summary>
<ul>
	<li>Parameters: <code>arg</code>, <code>func</code></li>
	<li>Parameter types: <code>Zahl</code>, <code>Buchstabe</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Führe die ANSI Funktion &lt;func&gt; mit dem argument &lt;arg&gt; aus&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe (ESC verkettet mit arg als Text verkettet mit func).

</code>
</pre>
</details>

<details>
<summary><h2>ANSI_Funktion_Ausführen</h2></summary>
<ul>
	<li>Parameters: <code>args</code>, <code>func</code></li>
	<li>Parameter types: <code>Zahlen Liste</code>, <code>Buchstabe</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Führe die ANSI Funktion &lt;func&gt; mit den argumenten &lt;args&gt; aus&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text befehl ist ESC verkettet mit args mit dem Trennzeichen ';' zum Text verbunden.
Schreibe (befehl verkettet mit func).

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Zeigen</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Zeige den Konsolen Zeiger&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe (ESC verkettet mit "?25h").

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Verstecken</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Verstecke den Konsolen Zeiger&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Schreibe (ESC verkettet mit "?25l").

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Hoch</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Bewege den Konsolen Zeiger um &lt;n&gt; Zeichen nach oben&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'A' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Runter</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Bewege den Konsolen Zeiger um &lt;n&gt; Zeichen nach unten&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'B' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Vorne</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Bewege den Konsolen Zeiger um &lt;n&gt; Zeichen nach vorne&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'C' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Hinten</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Bewege den Konsolen Zeiger um &lt;n&gt; Zeichen nach hinten&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'D' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Vorne_Zeile</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Bewege den Konsolen Zeiger um &lt;n&gt; Zeilen nach vorne&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'E' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Hinten_Zeile</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Bewege den Konsolen Zeiger um &lt;n&gt; Zeilen nach hinten&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'F' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Spalte</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Bewege den Konsolen Zeiger zu Spalte &lt;n&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'G' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeiger_Position</h2></summary>
<ul>
	<li>Parameters: <code>n</code>, <code>m</code></li>
	<li>Parameter types: <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Bewege den Konsolen Zeiger zu Zeile &lt;n&gt; und Spalte &lt;m&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'H' mit den argumenten (eine Liste, die aus n, m besteht) aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Leeren_Ende</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Leere die Konsole vom Zeiger zum Ende des Bildschirms&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'J' mit dem argument 0 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Leeren_Anfang</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Leere die Konsole vom Zeiger zum Anfang des Bildschirms&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'J' mit dem argument 1 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Leeren</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Leere die gesamte Konsole&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'J' mit dem argument 2 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Leeren_ClearBuf</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche alle gespeicherten Zeilen&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'J' mit dem argument 3 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeile_Leeren_Ende</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Leere die Zeile vom Zeiger zum Ende des Bildschirms&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'K' mit dem argument 0 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeile_Leeren_Anfang</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Leere die Zeile vom Zeiger zum Anfang des Bildschirms&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'K' mit dem argument 1 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Zeile_Leeren</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Leere die gesamte Zeile&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'K' mit dem argument 2 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Scrollen_Hoch</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Scrolle die Konsole um &lt;n&gt; Zeilen nach oben&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'S' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Scrollen_Runter</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Scrolle die Konsole um &lt;n&gt; Zeilen nach unten&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'T' mit dem argument n aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Schrift_Reset</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Konsolenschrift zurücksetzen&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 0 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Schrift_Fett</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Konsolenschrift fett setzen&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 1 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Schrift_Schräg</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Konsolenschrift schräg setzen&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 3 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Schrift_Unterstrichen</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Konsolenschrift unterschrichen setzen&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 4 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Schrift_Durchgestrichen</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Konsolenschrift durchgestrichen setzen&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 9 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Reset</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zurück.&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 39 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Schwarz</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zu schwarz&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 30 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Rot</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zu rot&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 31 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Grün</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zu grün&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 32 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Gelb</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zu gelb&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 33 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Blau</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zu blau&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 34 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Magenta</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zu magenta&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 35 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Türkis</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zu türkis&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 36 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_Weiß</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zu weiß&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 37 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_8Bit</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zum index &lt;n&gt; einer 8Bit Farbpallete&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit den argumenten (eine Liste, die aus 38, 5, n besteht) aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Vordergrund_RGB</h2></summary>
<ul>
	<li>Parameters: <code>r</code>, <code>g</code>, <code>b</code></li>
	<li>Parameter types: <code>Zahl</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift zum RGB wert &lt;r&gt;, &lt;g&gt;, &lt;b&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit den argumenten (eine Liste, die aus 38, 2, r, g, b besteht) aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Reset</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zurück.&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 49 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Schwarz</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zu schwarz&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 40 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Rot</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zu rot&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 41 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Grün</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zu grün&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 42 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Gelb</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zu gelb&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 43 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Blau</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zu blau&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 44 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Magenta</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zu magenta&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 45 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Türkis</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zu türkis&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 46 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_Weiß</h2></summary>
<ul>
</li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zu weiß&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit dem argument 47 aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_8Bit</h2></summary>
<ul>
	<li>Parameters: <code>n</code></li>
	<li>Parameter type: <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zum index &lt;n&gt; einer 8Bit Farbpallete&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit den argumenten (eine Liste, die aus 48, 5, n besteht) aus.

</code>
</pre>
</details>

<details>
<summary><h2>Konsole_Hintergrund_RGB</h2></summary>
<ul>
	<li>Parameters: <code>r</code>, <code>g</code>, <code>b</code></li>
	<li>Parameter types: <code>Zahl</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Konsolenschrift Hintergrund zum RGB wert &lt;r&gt;, &lt;g&gt;, &lt;b&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Führe die ANSI Funktion 'm' mit den argumenten (eine Liste, die aus 48, 2, r, g, b besteht) aus.

</code>
</pre>
</details>


