// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = fasthttp.Client{}
)

func Register(r chi.Router, s Server, opts ...Option) {
	r.MethodFunc("POST", "/addStickerToSet", NewAddStickerToSetHandler(s, opts...))
	r.MethodFunc("POST", "/answerCallbackQuery", NewAnswerCallbackQueryHandler(s, opts...))
	r.MethodFunc("POST", "/answerInlineQuery", NewAnswerInlineQueryHandler(s, opts...))
	r.MethodFunc("POST", "/answerPreCheckoutQuery", NewAnswerPreCheckoutQueryHandler(s, opts...))
	r.MethodFunc("POST", "/answerShippingQuery", NewAnswerShippingQueryHandler(s, opts...))
	r.MethodFunc("POST", "/banChatMember", NewBanChatMemberHandler(s, opts...))
	r.MethodFunc("POST", "/copyMessage", NewCopyMessageHandler(s, opts...))
	r.MethodFunc("POST", "/createChatInviteLink", NewCreateChatInviteLinkHandler(s, opts...))
	r.MethodFunc("POST", "/createNewStickerSet", NewCreateNewStickerSetHandler(s, opts...))
	r.MethodFunc("POST", "/deleteChatPhoto", NewDeleteChatPhotoHandler(s, opts...))
	r.MethodFunc("POST", "/deleteChatStickerSet", NewDeleteChatStickerSetHandler(s, opts...))
	r.MethodFunc("POST", "/deleteMessage", NewDeleteMessageHandler(s, opts...))
	r.MethodFunc("POST", "/deleteMyCommands", NewDeleteMyCommandsHandler(s, opts...))
	r.MethodFunc("POST", "/deleteStickerFromSet", NewDeleteStickerFromSetHandler(s, opts...))
	r.MethodFunc("POST", "/deleteWebhook", NewDeleteWebhookHandler(s, opts...))
	r.MethodFunc("POST", "/editChatInviteLink", NewEditChatInviteLinkHandler(s, opts...))
	r.MethodFunc("POST", "/editMessageCaption", NewEditMessageCaptionHandler(s, opts...))
	r.MethodFunc("POST", "/editMessageLiveLocation", NewEditMessageLiveLocationHandler(s, opts...))
	r.MethodFunc("POST", "/editMessageMedia", NewEditMessageMediaHandler(s, opts...))
	r.MethodFunc("POST", "/editMessageReplyMarkup", NewEditMessageReplyMarkupHandler(s, opts...))
	r.MethodFunc("POST", "/editMessageText", NewEditMessageTextHandler(s, opts...))
	r.MethodFunc("POST", "/exportChatInviteLink", NewExportChatInviteLinkHandler(s, opts...))
	r.MethodFunc("POST", "/forwardMessage", NewForwardMessageHandler(s, opts...))
	r.MethodFunc("POST", "/getChat", NewGetChatHandler(s, opts...))
	r.MethodFunc("POST", "/getChatAdministrators", NewGetChatAdministratorsHandler(s, opts...))
	r.MethodFunc("POST", "/getChatMember", NewGetChatMemberHandler(s, opts...))
	r.MethodFunc("POST", "/getChatMemberCount", NewGetChatMemberCountHandler(s, opts...))
	r.MethodFunc("POST", "/getFile", NewGetFileHandler(s, opts...))
	r.MethodFunc("POST", "/getGameHighScores", NewGetGameHighScoresHandler(s, opts...))
	r.MethodFunc("POST", "/getMe", NewGetMeHandler(s, opts...))
	r.MethodFunc("POST", "/getMyCommands", NewGetMyCommandsHandler(s, opts...))
	r.MethodFunc("POST", "/getStickerSet", NewGetStickerSetHandler(s, opts...))
	r.MethodFunc("POST", "/getUpdates", NewGetUpdatesHandler(s, opts...))
	r.MethodFunc("POST", "/getUserProfilePhotos", NewGetUserProfilePhotosHandler(s, opts...))
	r.MethodFunc("POST", "/leaveChat", NewLeaveChatHandler(s, opts...))
	r.MethodFunc("POST", "/pinChatMessage", NewPinChatMessageHandler(s, opts...))
	r.MethodFunc("POST", "/promoteChatMember", NewPromoteChatMemberHandler(s, opts...))
	r.MethodFunc("POST", "/restrictChatMember", NewRestrictChatMemberHandler(s, opts...))
	r.MethodFunc("POST", "/revokeChatInviteLink", NewRevokeChatInviteLinkHandler(s, opts...))
	r.MethodFunc("POST", "/sendAnimation", NewSendAnimationHandler(s, opts...))
	r.MethodFunc("POST", "/sendAudio", NewSendAudioHandler(s, opts...))
	r.MethodFunc("POST", "/sendChatAction", NewSendChatActionHandler(s, opts...))
	r.MethodFunc("POST", "/sendContact", NewSendContactHandler(s, opts...))
	r.MethodFunc("POST", "/sendDice", NewSendDiceHandler(s, opts...))
	r.MethodFunc("POST", "/sendDocument", NewSendDocumentHandler(s, opts...))
	r.MethodFunc("POST", "/sendGame", NewSendGameHandler(s, opts...))
	r.MethodFunc("POST", "/sendInvoice", NewSendInvoiceHandler(s, opts...))
	r.MethodFunc("POST", "/sendLocation", NewSendLocationHandler(s, opts...))
	r.MethodFunc("POST", "/sendMediaGroup", NewSendMediaGroupHandler(s, opts...))
	r.MethodFunc("POST", "/sendMessage", NewSendMessageHandler(s, opts...))
	r.MethodFunc("POST", "/sendPhoto", NewSendPhotoHandler(s, opts...))
	r.MethodFunc("POST", "/sendPoll", NewSendPollHandler(s, opts...))
	r.MethodFunc("POST", "/sendSticker", NewSendStickerHandler(s, opts...))
	r.MethodFunc("POST", "/sendVenue", NewSendVenueHandler(s, opts...))
	r.MethodFunc("POST", "/sendVideo", NewSendVideoHandler(s, opts...))
	r.MethodFunc("POST", "/sendVideoNote", NewSendVideoNoteHandler(s, opts...))
	r.MethodFunc("POST", "/sendVoice", NewSendVoiceHandler(s, opts...))
	r.MethodFunc("POST", "/setChatAdministratorCustomTitle", NewSetChatAdministratorCustomTitleHandler(s, opts...))
	r.MethodFunc("POST", "/setChatDescription", NewSetChatDescriptionHandler(s, opts...))
	r.MethodFunc("POST", "/setChatPermissions", NewSetChatPermissionsHandler(s, opts...))
	r.MethodFunc("POST", "/setChatPhoto", NewSetChatPhotoHandler(s, opts...))
	r.MethodFunc("POST", "/setChatStickerSet", NewSetChatStickerSetHandler(s, opts...))
	r.MethodFunc("POST", "/setChatTitle", NewSetChatTitleHandler(s, opts...))
	r.MethodFunc("POST", "/setGameScore", NewSetGameScoreHandler(s, opts...))
	r.MethodFunc("POST", "/setMyCommands", NewSetMyCommandsHandler(s, opts...))
	r.MethodFunc("POST", "/setPassportDataErrors", NewSetPassportDataErrorsHandler(s, opts...))
	r.MethodFunc("POST", "/setStickerPositionInSet", NewSetStickerPositionInSetHandler(s, opts...))
	r.MethodFunc("POST", "/setStickerSetThumb", NewSetStickerSetThumbHandler(s, opts...))
	r.MethodFunc("POST", "/setWebhook", NewSetWebhookHandler(s, opts...))
	r.MethodFunc("POST", "/stopMessageLiveLocation", NewStopMessageLiveLocationHandler(s, opts...))
	r.MethodFunc("POST", "/stopPoll", NewStopPollHandler(s, opts...))
	r.MethodFunc("POST", "/unbanChatMember", NewUnbanChatMemberHandler(s, opts...))
	r.MethodFunc("POST", "/unpinAllChatMessages", NewUnpinAllChatMessagesHandler(s, opts...))
	r.MethodFunc("POST", "/unpinChatMessage", NewUnpinChatMessageHandler(s, opts...))
	r.MethodFunc("POST", "/uploadStickerFile", NewUploadStickerFileHandler(s, opts...))
}
