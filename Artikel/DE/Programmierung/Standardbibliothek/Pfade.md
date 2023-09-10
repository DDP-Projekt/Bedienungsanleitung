# Duden/Pfade Funktionen
<details>
<summary><h2>TrennZeichen</h2></summary>
<ul>
<pre>
Gibt den Pfad-Trennzeichen, der dem Betriebssystem entspricht zurück.
Auf Linux: '/'
Auf Windows: '\'
</pre>
</li>
	<li>Rückgabe Typ: <code>Buchstabe</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Pfad-Trennzeichen&#34;</code></li>
	<li><code>&#34;den Pfad-Trennzeichen&#34;</code></li>
	<li><code>&#34;das Pfad-Trennzeichen&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn das Betriebssystem gleich "Windows" ist, gib '\\' zurück.
Gib '/' zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstAbsolut</h2></summary>
<ul>
<pre>
Gibt zurück ob der gegebene UNIX Pfad ein absoluter oder relativer Pfad ist
</pre>
	<li>Parameter: <code>t</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;t&gt; ein Absoluter Pfad ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib '/' am Anfang von t steht zurück.

</code>
</pre>
</details>

<details>
<summary><h2>OrdnerPfad</h2></summary>
<ul>
<pre>
Gibt den Pfad ohne den letzen Element zurück. 
Falls der Pfad leer ist, wird ein "." zurück gegeben.
Falls der Pfad nur aus "/" gefolg von nicht-"/" Zeichen, wird "/" zurückgegeben.
</pre>
	<li>Parameter: <code>t</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;den Ordnerpfad von &lt;t&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl i ist die Länge von t.

Solange i größer als 0 ist, mache:
	Wenn t an der Stelle i gleich '/' ist, dann:
		Der Text dir ist t von 1 bis i.
		Entferne alle '/' nach dir.
		Wenn die Länge von dir gleich 0 ist, gib "/" zurück.

		Gib dir zurück.
	Verringere i um 1.
Gib "." zurück.

</code>
</pre>
</details>

<details>
<summary><h2>BasisPfad</h2></summary>
<ul>
<pre>
Gibt den letzten Element eines Pfades zurück. Falls der Pfad leer ist, wird "." ausgegeben.
</pre>
	<li>Parameter: <code>t</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;den Basispfad von &lt;t&gt;&#34;</code></li>
	<li><code>&#34;der Basispfad von &lt;t&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text t2 ist (t mit allen '/' danach entfernt).
Die Zahl i ist die Länge von t2.
Solange i größer als 0 ist, mache:
	Wenn t2 an der Stelle i gleich '/' ist, dann:
		Gib t2 von (i plus 1) bis (die Länge von t2) zurück.
	Verringere i um 1.
Gib "." zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Erweiterung</h2></summary>
<ul>
<pre>
Gibt die Erweiterung einer Datei zurück. Wenn der Pfad eines Ordners gegeben wurde 
oder der Parameter t leer ist, wird einen leeren Text zurück gegeben.
</pre>
	<li>Parameter: <code>t</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Erweiterung der Datei bei &lt;t&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl i ist die Länge von t.
Solange i größer als 0 ist, mache:
	Wenn t an der Stelle i gleich '/' ist, gib "" zurück.
	Wenn t an der Stelle i gleich '.' ist, gib t von i bis (die Länge von t) zurück.
	Verringere i um 1.
Gib "" zurück.

</code>
</pre>
</details>

<details>
<summary><h2>DateiName</h2></summary>
<ul>
<pre>
Gibt den Dateinamen einer Datei zurück.
</pre>
	<li>Parameter: <code>t</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;den Dateinamen von &lt;t&gt;&#34;</code></li>
	<li><code>&#34;der Dateiname von &lt;t&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text bf ist der Basispfad von t.
Wenn bf gleich "." ist, gib "" zurück.
Die Zahl i ist 1.
Solange i kleiner als, oder die Länge von bf ist und bf an der Stelle i ungleich '.' ist, erhöhe i um 1.
Gib bf von 1 bis (i minus 1) zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Säubern</h2></summary>
<ul>
<pre>
Säubert/Normalisiert einen UNIX Pfad:
- Ersetzt mehrere aufeinderfolgende '/' durch einen ("///" -> "/")
- Entfernt '.' pfade ("a/." -> "a")
- Entfernt innere '..' pfade und das (nicht-..) element das davor steht ("a/b/.." -> "a")
- Entfernt .. elemente die nach root stehen ("/.." -> "/")

Der gesäuberter Pfad endet mit einem '/' nur wenn es root ist.
Ein leerer Pfad gibt "." zurück.

Entspricht Go's path.clean funktion
https://pkg.go.dev/path#Clean
</pre>
	<li>Parameter: <code>t</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Pfad &lt;t&gt; gesäubert&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn t gleich "" ist, gib "." zurück.
Der Boolean rooted ist wahr, wenn t an der Stelle 1 gleich '/' ist.
Die Zahl n ist die Länge von t.
Die Text path ist "". [ output text ]
Die Zahl r ist 1. [ next index to process ]
Die Zahl dotdot ist 1. [ index where .. must stop ]
Wenn rooted gleich wahr ist, dann:
	Füge '/' an path an.
	Speichere 2 in r.
	Speichere 2 in dotdot.

Solange r kleiner als, oder n ist, mache:
	[Schreibe r auf eine Zeile.]
	Wenn t an der Stelle r gleich '/' ist, dann:
		[ empty path element ]
		Erhöhe r um 1. [skip /]
	Wenn aber t an der Stelle r gleich '.' ist und (r gleich n ist oder t an der Stelle (r plus 1) gleich '/' ist), dann:
		[ . element ]
		Erhöhe r um 1. [skip .]
	Wenn aber t an der Stelle r gleich '.' ist und t an der Stelle (r plus 1) gleich '.' ist und (r plus 1 gleich n ist oder t an der Stelle (r plus 2) gleich '/' ist), dann:
		[ .. element: remove last / ]
		Erhöhe r um 2. 
		Wenn die Länge von path größer als, oder dotdot ist, dann:
			Mache:
				Wenn die Länge von path gleich 1 ist, dann:
					Speichere "" in path.
				Sonst
					Speichere path von 1 bis (die Länge von path minus 1) in path.
			Solange die Länge von path größer als dotdot ist und path an der Stelle (die Länge von path minus 1) ungleich '/' ist.
		Wenn aber rooted gleich falsch ist, dann:
			[ cannot backtrack, but not rooted, so append .. element ]
			Wenn die Länge von path größer als 0 ist, dann:
				Füge '/' an path an.
			Füge ".." an path an.
			Speichere die Länge von path in dotdot.
	Sonst:
		[ add slash if needed ]
		Wenn rooted und die Länge von path ungleich 1 ist oder nicht rooted und die Länge von path ungleich 0 ist, dann:
			Füge '/' an path an.

		Solange r kleiner als, oder n ist und t an der Stelle r ungleich '/' ist, mache:
			Füge (t an der Stelle r) an path an.
			Erhöhe r um 1.

Wenn path leer ist, speichere "." in path.
Gib path zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Verbinden</h2></summary>
<ul>
<pre>
Verbindet zwei UNIX Pfade mit einem '/' und säubert zuletzt.
</pre>
	<li>Parameter: <code>a</code>, <code>b</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Pfade &lt;a&gt; und &lt;b&gt; verbunden&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn a leer ist, gib b zurück.
Wenn b leer ist, gib a zurück.
Gib der Pfad (a verkettet mit '/' verkettet mit b) gesäubert zurück.

</code>
</pre>
</details>


