package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Section struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content template.HTML      `json:"content,omitempty" bson:"content,omitempty"`
}

type Sections []Section

var sections Sections
var mySections *mongo.Collection
var ctx context.Context

/*
*
Die main-Methode stellt die Verbindung mit der Datenbank her und startet die Webseiten auf dem Port:8080
*/
func main() {

	//Verbindung mit Datenbank
	ctx = context.Background()
	opt := options.Client().ApplyURI("mongodb://root:rootpassword@gomdb:27017")
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	// ueberpruefen der Verbindung
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	mySections = client.Database("mydb").Collection("Sections")
	sections, err = SectionInit(ctx, mySections)
	if err != nil {
		log.Fatal(err)
	}

	//REST Verbindung durch das Gin-Framework
	router := gin.Default()

	//aufruf der statischen Webseite
	router.Static("/static", "./static")

	//aufruf der template-Webseite
	router.LoadHTMLGlob("templates/**")
	router.GET("/template", pageHandler)

	//Verbindung auf Port:8080
	router.Run(":8080")
}

/*
*
Die Methode SectionInit speichert den HTML-Code in der Datenbank
*/
func SectionInit(ctx context.Context, coll *mongo.Collection) (Sections, error) {
	var a = []interface{}{
		bson.D{{Key: "title", Value: "About"},
			{Key: "content", Value: `<!-- MAIN -->
		<main>
		<!-- Unterüberschrift -->
		<h2 id="mich">Programmierer, Designer, Fotograf...</h2>
	   
		<!--PORTFOLIO-->
		<section class="portfolio">

			<div class="tile1">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/bild_2.jpg">
				</div>
			</div>

			<div class="tile1">
				<div class="text">
					<p>
					 Mein Name ist Adrian Busche und ich studiere zurzeit Medieninformatik an der Hochschule Harz in Wernigerode. 
					 In diesem Studiengang lerne ich nicht nur verschiedene Programmiersprachen, sondern auch diverse designtechnische Dinge. 
					 Dazu gehören unter anderem Grafikdesign, Fotografie, Filmtechnik und -schnitt oder auch Sounddesign. 
					 In diesem Portfolio zeige ich ein paar meiner Kreationen die teilweise im Zuge meines Studiums, aber auch außerhalb der Studienzeit entstanden sind.
					</p>
					
				</div>
			</div>

			<!-- Bild -->
			<div class="tile1">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/bild_1.jpg">
				</div>
			</div>

		</section><!--PORTFOLIO-->

		<!-- Map -->
		<div id="mapdiv"></div>
		
		<!-- Text -->
		<div class="text">
			<p>Mein Zuhause</p>
		</div>

		<!-- Unterüberschrift -->
		<h2 id="Fotografie">Fotografie</h2>
		
		<!-- Bild -->
			<div class="tile" >
				<div class="picture">
					<img class="lazy" data-src="../static/assets/img/bild_grashuepfer.jpg">
				</div>
			</div>
		   
		<!-- PORTFOLIO -->
		<section class="portfolio">

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/bild_3.jpg">
				</div>
			</div>

			<!-- Text -->
			<div class="tile">
				<div class="text">
				 Im bereich der Fotografie habe ich vor meinem Studium nicht viele Erfahrungen gesammelt, da ich nie eine gute Kamera besessen habe. 
				 Jedoch hat ich im laufe der Zeit bei mir ein Interesse daran entwickelt weswegen ich immer öfter meine Handykamera hervor hole, wenn ich ein Interessantes Motiv sehe.
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/bild_4.jpg">
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/bass_1.jpg">
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/bass_2.jpg">
					<p>Fotoserie einer Bassguitarre</p>
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/bass_3.jpg">
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/london_1.jpg">
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/london_2.jpg">
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/london_3.jpg">
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/bild.jpg">
				</div>
			</div>

			<!-- Text -->
			<div class="tile">
				<div class="text">
					In den meisten Fällen mache ich Landschafts- oder Naturfotos aber ich fotografiere auch gerne verschiedenste Objekte.
					Natürlich Fotografiere ich auch wie die meisten auch viel im Urlaub. Nur Portraits mache ich eher selten.
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/london_0.jpg">
				</div>
			</div>
		</section> <!-- PORTFOLIO -->

		<h2 id="Mediengestaltung">Mediengestaltung</h2> <!-- Unterüberschrift -->

		<!-- PORTFOLIO -->
		<section class="portfolio" >

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/Plakat_1.jpg">
				</div>
			</div>

			<!-- Text -->
			<div class="tile">
				<div class="text">
					<p>Im Fach Mediengestaltung habe ich einiges über die Grundlagen des Designens gelernt.
						Im Zuge dieses Kurses habe ich unter anderem Filmplakate und ein Coorperate Design eines Weingutes entwurfen.
					</p>
				</div>
			</div>

			<!-- Bild -->
			<div class="tile">
				<div class="pic">
					<img class="lazy" data-src="../static/assets/img/Plakat_2.jpg">
				</div>
			</div>
		</section> <!-- PORTFOLIO -->

		<!-- Bild -->
		<div class="tile">
			<div class="picture">
				<img class="lazy" data-src="../static/assets/img/Weinflaschen.jpg">
			</div>
		</div>
		
		<h2 id="Sounddesign">Sounddesign</h2> <!-- Unterüberschrift -->

		<!-- Portfolio -->
		<section class="portfolio" >

			<!-- Video -->
			<div class="tile">
				<div class="video">
					<video id="video" controls preload="metadata">
						<source src="../static/assets/video/ufo.mp4" type="video/mp4">
					</video>
				</div>
			</div>

			<!-- Text -->
			<div class="tile">
				<div class="text">
					<p>
						Auch im bereich Sounddesign habe ich in meinem Studium vieles gelernt.
						Zum einen habe ich durch Sounds eine Alienentführung dargestellt und zum anderen habe ich ein eingesprochenes Gedicht mit Geräuschen untermahlt.
					</p>
				</div>
			</div>

			<!-- Video -->
			<div class="tile">
				<div class="video">
					<video id="video" controls preload="metadata">
						<source src="../static/assets/video/gedicht.mp4" type="video/mp4">
					</video>
				</div>
			</div>
		</section> <!-- PORTFOLIO -->

		<h2 id="Sonstiges">Sonstiges</h2> <!-- Unterüberschrift -->

		<!-- PORTFOILIO -->
		<section class="portfolio" >

			<!-- Text -->
			<div class="tile">
				<div class="text">
					Ich habe mich, ausserhalb meines Studiums, vor einiger Zeit auch einmal daran versucht Clips die ich beim Zocken Aufgenommen habe, zu schneiden und zu bearbeiten.
				</div>
			</div>

			<!-- Video -->
			<div class="tile">
				<div class="video">
					<iframe width="100%" height="auto" src="https://www.youtube-nocookie.com/embed/healHY7HGew"
					frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture"
					allowfullscreen>
					</iframe>
				</div>
			</div>
		</section> <!-- PORTFOLIO -->

	</main>
	<!-- MAIN END -->`},
		},
	}

	if err := coll.Drop(ctx); err != nil {
		log.Printf("db not dropped %v", err)
	}

	_, err := mySections.InsertMany(ctx, a)
	if err != nil {
		return nil, err
	}

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var all Sections
	if err = cursor.All(ctx, &all); err != nil {
		return nil, err
	}

	return all, nil
}

/*
*
Die Methode pageHandler laedt den HTML-Quellcode aus der Datenbank
*/
func pageHandler(c *gin.Context) {
	cursor, err := mySections.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var all Sections
	if err = cursor.All(ctx, &all); err != nil {
		return
	}

	c.HTML(http.StatusOK, "index", gin.H{"allSections": all})
}
