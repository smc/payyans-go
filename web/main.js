function updateMap() {
  var mapUrl = $('#mapselect').val();
  mapUrl = './font-maps/' + mapUrl + '.map';
  $.ajax({
    url: mapUrl,
    success: function(data) {
      $('#map').val(data);
    },
    dataType: 'text'
  });
}

function updateNormalizerRules() {
  mapUrl = './normalizer-rules/normalizer_ml.rules';
  $.ajax({
    url: mapUrl,
    success: function(data) {
      $('#normalizer-rules').val(data);
    },
    dataType: 'text'
  });
}

document.querySelector('#mapselect').addEventListener('change', function(e){
  updateMap()
});

document.querySelector('#convert-ascii').addEventListener('click', function(e){
  document.querySelector('#unicode-input').value = "Converting...";
  
  const fontMap = document.querySelector('#map').value;
  const normalizerRules = document.querySelector('#normalizer-rules').value;

  const result = AsciiToUnicodeByMapString(document.querySelector('#ascii-input').value, fontMap, normalizerRules);
  
  document.querySelector('#unicode-input').value = result
});

updateMap()
updateNormalizerRules()
