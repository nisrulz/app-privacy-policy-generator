function fc_deploy_simple() {
  var form = document.getElementById("fc-form");

  // fill in the body with the html content
  document.getElementById('fc-body').value = getContent('privacy_simple_content')
  document.getElementById('fc-terms').value = getContent('tandc_content')
  
  form.submit();
}

function fc_deploy_notracking() {
  var form = document.getElementById("fc-form");

  // fill in the body with the html content
  document.getElementById('fc-body').value = getContent('privacy_notrack_content')
  document.getElementById('fc-terms').value = getContent('tandc_content')
  
  form.submit();
}

function fc_deploy_gdpr() {
  var form = document.getElementById("fc-form");

  // fill in the body with the html content
  document.getElementById('fc-body').value = getContent('privacy_gdpr_content')
  document.getElementById('fc-terms').value = getContent('tandc_content')
  
  form.submit();
}