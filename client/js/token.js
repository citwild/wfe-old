/**
 * Created by milesdowe on 5/7/16.
 */

/**
 * For setting the JWT in the sessionStorage
 * @param token
 */
function setJwtToken(token) {
    sessionStorage.setItem('token', token);
}
