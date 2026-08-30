window.fc_deploy = function (bodyContentId) {
  var form = document.getElementById("fc-form");

  // fill in the body with the html content
  document.getElementById('fc-body').value = getContent(bodyContentId)
  document.getElementById('fc-terms').value = getContent('tandc_content')

  form.submit();
}
