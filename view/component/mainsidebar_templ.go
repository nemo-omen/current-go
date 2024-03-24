// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package component

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"newser.app/internal/dto"
	"strconv"
)

func getSidebarFeedInfo(ctx context.Context) []*dto.FeedInfoDTO {
	feedInfos := ctx.Value("feedlinks")
	if feedInfos != nil {
		return feedInfos.([]*dto.FeedInfoDTO)
	}
	return []*dto.FeedInfoDTO{}
}

func getFeedLinkCount(nf *dto.FeedInfoDTO) string {
	return strconv.Itoa(nf.UnreadCount)
}

func getCurrentPath(ctx context.Context) string {
	currentPath := ctx.Value("currentPath")
	if currentPath != nil {
		return currentPath.(string)
	}
	return ""
}

func isCurrentString(ctx context.Context, path string) string {
	// fmt.Println("isCurrentString", path, getCurrentPath(ctx))
	if getCurrentPath(ctx) == path {
		return "page"
	}
	return ""
}

func MainSidebar() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<aside class=\"sidebar\" id=\"sidebar-main\"><nav aria-label=\"Secondary Navigation\" hx-boost=\"true\" hx-disinherit=\"*\"><ul><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = IconLink(
			"/app",
			"list",
			"All Posts",
			templ.Attributes{
				"class":        "icon-link",
				"aria-current": isCurrentString(ctx, "/app"),
				"hx-get":       "/app",
				"hx-target":    "main",
				"hx-swap":      "innerHTML",
				"hx-push-url":  "true",
			},
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = IconLink(
			"/app/collection/saved",
			"bookmark",
			"Saved",
			templ.Attributes{
				"class":        "icon-link",
				"aria-current": isCurrentString(ctx, "/app/collection/saved"),
				"hx-get":       "/app/collection/saved",
				"hx-target":    "main",
				"hx-swap":      "innerHTML",
				"hx-push-url":  "true",
			},
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = IconLink(
			"/app/notes",
			"note",
			"Notes",
			templ.Attributes{
				"class":        "icon-link",
				"aria-current": isCurrentString(ctx, "/app/notes"),
				"hx-get":       "/app/note",
				"hx-target":    "main",
				"hx-swap":      "innerHTML",
				"hx-push-url":  "true",
			},
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = IconLink(
			"/app/search",
			"folder_add",
			"Add Feed",
			templ.Attributes{
				"class":        "icon-link",
				"aria-current": isCurrentString(ctx, "/app/search"),
				"hx-get":       "/app/search",
				"hx-target":    "main",
				"hx-swap":      "innerHTML",
				"hx-push-url":  "true",
			},
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li></ul></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(getSidebarFeedInfo(ctx)) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"nav-vertical\" aria-label=\"Subscriptions\" hx-get=\"/app/control/unreadcount\" hx-trigger=\"updateUnreadCount from:body\" hx-target=\"#main-feed-links\" hx-swap=\"outerHTML\" hx-push-url=\"false\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = MainFeedLinks().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</nav>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</aside>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MainFeedLinks() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"main-feed-links\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, feedInfo := range getSidebarFeedInfo(ctx) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a class=\"icon-link\" href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 templ.SafeURL = templ.SafeURL("/app/newsfeed/" + feedInfo.FeedId)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" aria-current=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(isCurrentString(ctx, "/app/newsfeed/"+feedInfo.FeedId)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/app/newsfeed/" + feedInfo.FeedId))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"main\" hx-swap=\"innerHTML\" hx-push-url=\"true\"><img src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(feedInfo.ImageUrl))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"feedLink.ImageTitle\" class=\"image-icon\"> <span class=\"link-text\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(feedInfo.FeedTitle)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/component/mainsidebar.templ`, Line: 151, Col: 26}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if feedInfo.UnreadCount > 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"badge\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(getFeedLinkCount(feedInfo))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/component/mainsidebar.templ`, Line: 155, Col: 35}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul><script>\n\t\tconst sidebar = document.getElementById('sidebar-main');\n\t\tconst links = sidebar.querySelectorAll('a');\n\t\thtmx.on('htmx:afterSettle', function(evt) {\n\t\t\tconst currentUrl = new URL(window.location);\n\t\t\tconst currentPath = currentUrl.pathname;\n\t\t\tlinks.forEach(link => {\n\t\t\t\tconst linkUrl = new URL(link.href);\n\t\t\t\tconst linkPath = linkUrl.pathname;\n\t\t\t\tif (currentPath === linkPath) {\n\t\t\t\t\tlink.setAttribute('aria-current', 'page');\n\t\t\t\t} else {\n\t\t\t\t\tlink.setAttribute('aria-current', '');\n\t\t\t\t}\n\t\t\t});\n\t\t});\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
