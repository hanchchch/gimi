const host = `https://some-evil-backend.${Math.round(
  Math.random() * 36 ** 8
).toString(36)}-hanchchch.com`;

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
