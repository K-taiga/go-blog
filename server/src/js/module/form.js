// 厳格なモード
"use strict";

// DOM Tree の構築が完了したら処理を開始
document.addEventListener("DOMContentLoaded", () => {
  // DOM API を利用して HTML 要素を取得
  const inputs = document.getElementsByTagName("input");
  const form = document.forms.namedItem("article-form");
  const saveBtn = document.querySelector(".article-form__save");
  const cancelBtn = document.querySelector(".article-form__cancel");
  const previewOpenBtn = document.querySelector(".article-form__open-preview");
  const previewCloseBtn = document.querySelector(
    ".article-form__close-preview"
  );
  const articleFormBody = document.querySelector(".article-form__body");
  const articleFormPreview = document.querySelector(".article-form__preview");
  const articleFormBodyTextArea = document.querySelector(
    ".article-form__input--body"
  );
  const articleFormPreviewTextArea = document.querySelector(
    ".article-form__preview-body-contents"
  );

  // 新規作成画面か編集画面かurlで判定するための構造体をセット
  const mode = { method: "", url: "" };
  if (window.location.pathname.endsWith("new")) {
    mode.method = "POST";
    mode.url = "/";
  } else if (window.location.pathname.endsWith("edit")) {
    mode.method = "PATCH";
    //'/'以降の/:articleIDを取得する
    mode.url = `/${window.location.pathname.split("/")[1]}`;
  }
  const { method, url } = mode;

  for (let elm of inputs) {
    elm.addEventListener("keydwon", (event) => {
      // Enterキーが押された場合
      if (event.keyCode && event.keyCode === 13) {
        // デフォルトの挙動をさせない
        event.preventDefault();

        // 何もせず処理を終了する
        return false;
      }
    });
  }

  // プレビューを開くイベント
  previewOpenBtn.addEventListener("click", (event) => {
    // formのTextAreaに入力された内容をコピー
    articleFormPreviewTextArea.innerHTML = articleFormBodyTextArea.value;

    // 入力フォームを非表示
    articleFormBody.style.display = "none";

    // プレビューの表示
    articleFormPreview.style.display = "grid";
  });

  // プレビューを閉じるイベント
  previewCloseBtn.addEventListener("click", (event) => {
    // 入力フォームを表示
    articleFormBody.style.display = "grid";

    // プレビューを非表示
    articleFormPreview.style.display = "none";
  });

  // 前のページに戻るイベント
  cancelBtn.addEventListener("click", (event) => {
    // buttonのデフォルトの挙動を制御
    event.preventDefault();

    // 指定のURLに遷移
    window.location.href = url;
  });
});
