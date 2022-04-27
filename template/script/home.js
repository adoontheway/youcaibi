$(document).ready(function() {
    DEFAULT_COOLIE_EXPIRE_TIME = 30;// in minutes
    uname = "";
    session = "";
    uid = 0;
    current_video = null;
    listed_videos = null;

    session = getCookie('session');
    uname = getCookie('username');

    // home page event registry
    $("#regbtn").on("click",function (e) {
        
    });

    $("#siginbtn").on("click",function (e) {
        
    });

    $("#siginhref").on("click",function (e) {
        $("#regsubmit").hide()
        $("#siginsubmit").show()
    });

    $("#registerhref").on("click",function (e) {
        $("#regsubmit").show()
        $("#siginsubmit").hide()
    });

    // userhome page event registry
    $("#uploadform").on("submit",function (e) {
       
    });

    $(".close").on("click",function (e) {
       
    });

    $("#logout").on("click",function (e) {
       
    });
});
// init page and make callback
function initPage(callback) {
    
}

function setCookie(key, value, exmin) {
    
}

function getCookie(key) {
    
}
// DOM operations
function selectVideo(vid) {
    
}

function refreshComments(vid) {
    
}

function popupNotificationMessage(msg) {
    
}

function popupErrorMessage(msg) {
    
}

function htmlCommentListElement(cid, author,content) {
    
}

function htmlVideoListElement(vid, name, ctime) {
    
}
// ajax calls
//user related
function registerUser(callback) {
    
}

function signinUser(callback) {
    
}

function getUserId(callback) {
    
}
// video related
function createVideo(vname, callback) {
    
}

function listAllVideos(callback) {
    
}

function deleteVideo(vid,callback) {
    
}
//comment related
function postComment(vid, content,callback) {
    
}

function listAllComments(vidmcallback) {
    
}