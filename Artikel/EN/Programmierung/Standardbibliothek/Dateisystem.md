# Duden/Dateisystem functions
<details>
<summary><h2>Schreibe_Text_Datei</h2></summary>
<ul>
<pre>
Die Funktion Schreibe_Text_Datei schreibt einen Text (text) in die Datei an dem gegebenen Text Pfad.
Falls möglich Fehler auftreten, werden diese in fehler gespeichert.
</pre>
	<li>Parameters: <code>Pfad</code>, <code>text</code>, <code>fehler</code></li>
	<li>Parameter types: <code>Text</code>, <code>Text</code>, <code>Text Referenz</code></li>
	<li>Return type: <code>Zahl</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Schreibe den Text &lt;text&gt; in die Datei &lt;Pfad&gt; und speichere einen möglichen Fehler in &lt;fehler&gt;&#34;</code></li>
	<li><code>&#34;Schreibe den Text &lt;text&gt; in die Datei &lt;Pfad&gt; und speichere einen moeglichen Fehler in &lt;fehler&gt;&#34;</code></li>
	<li><code>&#34;die Anzahl der Bytes, die von &lt;text&gt; in &lt;Pfad&gt; geschrieben wurden, wobei ein möglicher Fehler in &lt;fehler&gt; gespeichert wurde&#34;</code></li>
	<li><code>&#34;die Anzahl der Bytes, die von &lt;text&gt; in &lt;Pfad&gt; geschrieben wurden, wobei ein moeglicher Fehler in &lt;fehler&gt; gespeichert wurde&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implemented in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Lies_Text_Datei</h2></summary>
<ul>
<pre>
Die Funktion Lies_Text_Datei speichert den Inhalt der Datei, die an dem gegebenen Pfad liegt, in ref und gibt die Anzahl der Bytes der gelesenen Datei zurück.
Wenn ein Fehler auftreten sollte, ist der zurückgegebene Wert negativ und die Fehler meldung in ref geschrieben
</pre>
	<li>Parameters: <code>Pfad</code>, <code>ref</code></li>
	<li>Parameter types: <code>Text</code>, <code>Text Referenz</code></li>
	<li>Return type: <code>Zahl</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lies den Text in &lt;Pfad&gt; und speichere ihn in &lt;ref&gt;&#34;</code></li>
	<li><code>&#34;die Anzahl der Bytes, die aus &lt;Pfad&gt; gelesen und in &lt;ref&gt; gespeichert wurden&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implemented in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Existiert_Pfad</h2></summary>
<ul>
<pre>
Überprüft ob der gegebene Pfad existiert (egal ob als Ordner oder Datei)
</pre>
	<li>Parameters: <code>Pfad</code></li>
	<li>Parameter type: <code>Text</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;die Datei &lt;Pfad&gt; existiert&#34;</code></li>
	<li><code>&#34;der Ordner &lt;Pfad&gt; existiert&#34;</code></li>
	<li><code>&#34;der Pfad &lt;Pfad&gt; existiert&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implemented in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Ist_Ordner</h2></summary>
<ul>
<pre>
Überprüft ob der gegebene Pfad ein Ordner ist
</pre>
	<li>Parameters: <code>Pfad</code></li>
	<li>Parameter type: <code>Text</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;&lt;Pfad&gt; ein Ordner ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implemented in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Ist_Datei</h2></summary>
<ul>
<pre>
Überprüft ob der gegebene Pfad eine Datei ist
</pre>
	<li>Parameters: <code>Pfad</code></li>
	<li>Parameter type: <code>Text</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;&lt;Pfad&gt; eine Datei ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von Pfad gleich 0 ist oder Pfad ein Ordner ist, gib falsch zurück.
Gib wahr zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Erstelle_Ordner</h2></summary>
<ul>
<pre>
Erstellt rekursiv den Ordner Pfad.
Rekursiv -> alle benötigten zwischen Ordner werden ebenfalls erstellt.

Gibt zurück ob das Erstellen erfolgreich war.
</pre>
	<li>Parameters: <code>Pfad</code></li>
	<li>Parameter type: <code>Text</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Erstelle den Ordner &lt;Pfad&gt;&#34;</code></li>
	<li><code>&#34;der Ordner &lt;Pfad&gt; erfolgreich erstellt wurde&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implemented in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Loesche_Pfad</h2></summary>
<ul>
<pre>
!!!Nicht unbedingt sicher!!!

Löscht die gegebene Datei oder den gegebenen Ordner.
Im Falle eines Ordners wird rekursiv das gesamte Verzeichnis gelöscht.

Gibt zurück ob das Löschen erfolgreich war.
</pre>
	<li>Parameters: <code>Pfad</code></li>
	<li>Parameter type: <code>Text</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche &lt;Pfad&gt;&#34;</code></li>
	<li><code>&#34;&lt;Pfad&gt; erfolgreich gelöscht wurde&#34;</code></li>
	<li><code>&#34;Lösche die Datei &lt;Pfad&gt;&#34;</code></li>
	<li><code>&#34;die Datei &lt;Pfad&gt; erfolgreich gelöscht wurde&#34;</code></li>
	<li><code>&#34;Lösche den Ordner &lt;Pfad&gt;&#34;</code></li>
	<li><code>&#34;der Ordner &lt;Pfad&gt; erfolgreich gelöscht wurde&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implemented in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>Pfad_Verschieben</h2></summary>
<ul>
<pre>
!!!Nicht unbedings sicher!!!

Verschiebt den Pfad zu NeuerName.
Kann auch zum Umbenennen benutzt werden.

Gibt zurück ob das Umbenennen erfolgreich war.
</pre>
	<li>Parameters: <code>Pfad</code>, <code>NeuerName</code></li>
	<li>Parameter types: <code>Text</code>, <code>Text</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Verschiebe &lt;Pfad&gt; nach &lt;NeuerName&gt;&#34;</code></li>
	<li><code>&#34;&lt;Pfad&gt; erfolgreich nach &lt;NeuerName&gt; verschoben wurde&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implemented in <code>"libddpstdlib.a"</code>
</details>


