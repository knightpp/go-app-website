package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/XANi/loremipsum"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type navbar struct {
	app.Compo
}

func (n *navbar) Render() app.UI {
	return app.Nav().Body(
		app.Div().Body(app.Img().Src("/web/img/tux-logo.svg").Height(50).Class("navbar-brand")),
		app.A().Href("https://vk.com").Text("Вконтакте").Class("nav-link"),
		app.A().Href("https://ok.ru").Text("Одноклассники").Class("nav-link"),
		app.A().Href("https://mail.ru").Text("Mail ru").Class("nav-link"),
		app.A().Href("https://twitter.com").Text("Twitter").Class("nav-link"),
	).Class("navbar", "navbar-expand-lg", "navbar-light", "bg-light", "justify-content-center")
}

func newFakeContent() *fakeContent {
	return &fakeContent{
		generator: loremipsum.New(),
	}
}

type fakeContent struct {
	app.Compo

	generator *loremipsum.LoremIpsum
}

func (c *fakeContent) Render() app.UI {
	return app.P().Text(c.generator.Paragraph())
}

type carouselItem struct {
	Image   string
	Caption string
}

var cn uint32

func newCarousel(items ...carouselItem) *carousel {
	id := atomic.AddUint32(&cn, 1)
	return &carousel{
		id:    "carousel-" + strconv.FormatUint(uint64(id), 10),
		Items: items,
	}
}

type carousel struct {
	app.Compo
	id    string
	Items []carouselItem
}

func (c *carousel) Render() app.UI {
	return app.Div().ID(c.id).Class("carousel", "slide").Attr("data-bs-ride", "carousel").Body(
		app.Div().Class("carousel-indicators").Body(
			app.Range(c.Items).Slice(func(i int) app.UI {
				button := app.Button().
					Attr("data-bs-target", "#"+c.id).
					Attr("data-bs-slide-to", strconv.Itoa(i))
				if i == 0 {
					button = button.Class("active").Attr("aria-current", "true")
				}
				return button
			}),
		),
		app.Div().Class("carousel-inner").Body(
			app.Range(c.Items).Slice(func(i int) app.UI {
				classes := []string{"carousel-item"}
				if i == 0 {
					classes = append(classes, "active")
				}
				return app.Div().Class(classes...).Body(
					app.Img().Src(c.Items[i].Image).Class("d-block w-100"),
					app.Div().Class(
						"carousel-caption",
						"d-none", "d-md-block",
						"bg-dark", "p-2", "bg-opacity-25").
						Body(
							app.P().Text(c.Items[i].Caption),
						),
				)
			}),
		),
		app.Button().Class("carousel-control-prev").
			Attr("data-bs-target", "#"+c.id).
			Attr("data-bs-slide", "prev").
			Body(
				app.Span().
					Class("carousel-control-prev-icon").
					Attr("aria-hidden", "true"),
				app.Span().Class("visually-hidden").Text("Previous"),
			),
		app.Button().Class("carousel-control-next").
			Attr("data-bs-target", "#"+c.id).
			Attr("data-bs-slide", "next").
			Body(
				app.Span().
					Class("carousel-control-next-icon").
					Attr("aria-hidden", "true"),
				app.Span().Class("visually-hidden").Text("Next"),
			),
	)
}

func newToastStack(n int) *toastStack {
	ts := toastStack{}
	for i := 0; i < n; i++ {
		ts.addAd()
	}
	return &ts
}

type toast struct {
	app.Compo
	Title string
	Msg   string
}

func (t *toast) OnMount(ctx app.Context) {
	ctx.NewActionWithValue("toast-added", t)
}

func (t *toast) Render() app.UI {
	return app.Div().Class("toast").
		Attr("role", "alert").
		Attr("aria-live", "assertive").
		Attr("aria-atomic", "true").
		Attr("data-bs-autohide", "false").
		Body(
			app.Div().Class("toast-header").Body(
				app.Strong().Class("me-auto").Text(t.Title),
				app.Button().Class("btn-close").
					Attr("data-bs-dismiss", "toast").
					Attr("aria-label", "Close"),
			),
			app.Div().Class("toast-body").Text(t.Msg),
		)
}

type toastStack struct {
	app.Compo
	Toasts []*toast
}

func (t *toastStack) OnMount(ctx app.Context) {
	ctx.Handle("toast-added", t.handleAddedToast)
}

func (t *toastStack) handleAddedToast(ctx app.Context, a app.Action) {
	toast := a.Value.(*toast)
	jsVal := toast.JSValue()
	jsVal.Set("show", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		app.Window().Get("bootstrap").Get("Toast").New(this).Call("show")
		return nil
	}))
	jsVal.Call("show")
	sleep := rand.Intn(15) + 3
	go func() {
		lorem := loremipsum.New()
		for {
			if jsVal.Get("classList").Call("contains", "hide").Bool() {
				toast.Msg = lorem.Sentence()
				toast.Title = lorem.Word()
				toast.Update()
				jsVal.Call("show")
			}
			time.Sleep(time.Duration(sleep) * time.Second)
		}
	}()
}
func (t *toastStack) addAd() {
	t.Toasts = append(t.Toasts, &toast{Title: "Реклама", Msg: "Новая реклама"})
}

func (t *toastStack) Render() app.UI {
	return app.Div().Class("toast-container", "position-fixed", "bottom-0", "end-0", "p-3").Body(
		app.Range(t.Toasts).Slice(func(i int) app.UI {
			return t.Toasts[i]
		}),
	)
}

type index struct {
	app.Compo
}

func (h *index) Render() app.UI {
	return app.Div().Body(
		&navbar{},
		app.Div().Class("container").Body(
			newFakeContent(),
		),
		app.Div().
			Class("container").
			Body(
				newCarousel(
					carouselItem{Image: "https://mir-s3-cdn-cf.behance.net/project_modules/1400/e166b544913357.5822024088cc1.jpg",
						Caption: "ВКОНТАКТЕ вредит здоровью"},
					carouselItem{Image: "https://www.lolwot.com/wp-content/uploads/2015/08/20-funny-and-creative-stock-images-found-online-5.jpg",
						Caption: "Однокласники признана худшей соц сетью"},
					carouselItem{Image: "https://www.evidentlycochrane.net/wp-content/uploads/2017/08/to-have-a-pineapple-picture-id108352271.jpg",
						Caption: "VK и mailru обьединяются в одну компанию"},
				),
			),
		newToastStack(4),
	)
}

func main() {
	// Components routing:
	app.Route("/", &index{})
	app.Route("/index", &index{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css",
		},
		Scripts: []string{
			"https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js",
		},
	})

	addr := "127.0.0.1:8010"
	log.Println(addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
