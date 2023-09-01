## Services

rm -rf account base cms community configuration finance hrm insight marketing payment sales
./init.sh account base cms community configuration finance hrm insight marketing payment sales

{
  aud: 'https://api.trip.express/graphql',
  azp: '64cc84de610349db6ca68102',
  exp: 1692698524,
  iat: 1692612124,
  iss: 'https://api.trip.express',
  sub: '64e07121ebebfd9aa840ae8c'
}

https://auth0.com/docs/secure/tokens/json-web-tokens/create-custom-claims

/**
 * Register any authentication / authorization services.
 */
public function boot(): void
{
    Passport::tokensExpireIn(now()->addDays(15));
    Passport::refreshTokensExpireIn(now()->addDays(30));
    Passport::personalAccessTokensExpireIn(now()->addMonths(6));
}

# Refreshing Tokens
 
$response = Http::asForm()->post('http://passport-app.test/oauth/token', [
    'grant_type' => 'refresh_token',
    'refresh_token' => 'the-refresh-token',
    'client_id' => 'client-id',
    'client_secret' => 'client-secret',
    'scope' => '',
]);
 

# Requesting Tokens
 
$response = Http::asForm()->post('http://passport-app.test/oauth/token', [
    'grant_type' => 'password',
    'client_id' => 'client-id',
    'client_secret' => 'client-secret',
    'username' => 'taylor@laravel.com',
    'password' => 'my-password',
    'scope' => '',
]);
 
