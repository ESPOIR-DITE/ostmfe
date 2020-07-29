let checkSelectedApplicantType = function(initialApplicantType) {
    const selectedApplicantType = $("#applicantType").val();
    $("#updateApplicantTypeBtn").hide();
    if (initialApplicantType.localeCompare(selectedApplicantType)) {
        $("#updateApplicantTypeBtn").show();
    }
};