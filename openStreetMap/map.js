const axios = require('axios');

const bounds = '35.682839,139.704218,35.707914,139.753842';

const overpassQuery = `
  [out:json];
  way
    [highway]
    (${bounds});
  out;
`;

axios.get(`https://overpass-api.de/api/interpreter?data=${encodeURIComponent(overpassQuery)}`)
  .then(response => {
    const data = response.data;

    if (data && data.elements) {
      const roads = data.elements.filter(element => element.type === 'way');
      axios.post('http://localhost:8002/post', { roads })
        .then(response => {
          console.log('データを正常に送信しました。');
        })
        .catch(error => {
          console.error('データの送信中にエラーが発生しました:', error);
        });
    } else {
      console.log('データが見つかりませんでした。');
    }
  })
  .catch(error => {
    console.error('データの取得中にエラーが発生しました:', error);
  });

