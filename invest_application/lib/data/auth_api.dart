import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import './config_loader.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:flutter/foundation.dart';
import 'package:shared_preferences/shared_preferences.dart';
// import 'package:http/browser_client.dart';



// JWT Token model
class AuthToken {
  final String token;
  final DateTime expiryDate;

  AuthToken({required this.token, required this.expiryDate});

  bool get isValid => DateTime.now().isBefore(expiryDate);
}

class AuthService {

  final String baseUrl = ConfigLoader.getUrl();
  final FlutterSecureStorage _secureStorage = const FlutterSecureStorage(); // in-device storage
  

  
  // Key for token storage
  static const String _tokenKey = 'jwt_token';
  static const String _tokenExpiryKey = 'jwt_token_expiry';
  
  // For notifying UI about auth state changes
  final ValueNotifier<bool> authStateChanges = ValueNotifier<bool>(false);
  
  // Singleton pattern
  static final AuthService _instance = AuthService._internal();
  
  factory AuthService() {
    return _instance;
  }
  
  AuthService._internal();
  
  // Register a new user
  Future<bool> register(String username, String email, String password) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/register'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode({
          'username': username,
          'email': email,
          'password': password,
        }),
      );
      
      if (response.statusCode == 201) {
        // Successfully registered, now login
        return login(email, password);
      }
      return false;
    } catch (e) {
      // print('Registration error: $e');
      return false;
    }
  }
  
  Future<bool> login(String user, String password) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/login'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode({
          'username': user,
          'password': password,
        }),
      );

      if (response.statusCode == 200) {
        final responseData = json.decode(response.body);
        
        // Parse token and expiry date
        final token = responseData['token'];
        final expiryTimeStamp = responseData['expiry']; // Unix timestamp
        final expiryDate = DateTime.fromMillisecondsSinceEpoch(expiryTimeStamp * 1000);
        
        // Save token based on platform
        await _saveToken(token, expiryDate);
        
        // Notify UI about auth state change
        authStateChanges.value = true;
        return true;
      } else {
        print('Login failed: ${response.statusCode} - ${response.body}');
        return false;
      }
    } catch (e) {
      print('Login exception: $e');
      return false;
    }
  }

  Future<void> _saveToken(String token, DateTime expiryDate) async {
    if (kIsWeb) {
      // For web platform, use shared_preferences
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString(_tokenKey, token);
      await prefs.setString(_tokenExpiryKey, expiryDate.millisecondsSinceEpoch.toString());
    } else {
      // For mobile platforms, use secure storage
      await _secureStorage.write(key: _tokenKey, value: token);
      await _secureStorage.write(key: _tokenExpiryKey, value: expiryDate.millisecondsSinceEpoch.toString());
    }
  }
  
  // Get token from secure storage
  Future<AuthToken?> getToken() async {
    dynamic token;
    dynamic expiryString;

    if (kIsWeb) {
      final prefs = await SharedPreferences.getInstance();
      token = prefs.getString(_tokenKey);
      expiryString =  prefs.getString(_tokenExpiryKey);
    } else {
    token = await _secureStorage.read(key: _tokenKey);
    expiryString = await _secureStorage.read(key: _tokenExpiryKey);
    
    }

    if (token != null && expiryString != null) {
      final expiryDate = DateTime.fromMillisecondsSinceEpoch(int.parse(expiryString));
      return AuthToken(token: token, expiryDate: expiryDate);
    }
    return null;
  }
  
  // Check if user is authenticated
  Future<bool> isAuthenticated() async {
    final token = await getToken();
    return token != null && token.isValid;
  }
  
  // Logout - remove token from storage
  Future<void> logout() async {
    if (kIsWeb) {
      final prefs = await SharedPreferences.getInstance();
      await prefs.remove(_tokenKey);
      await prefs.remove(_tokenExpiryKey);
    } else {
      await _secureStorage.delete(key: _tokenKey);
      await _secureStorage.delete(key: _tokenExpiryKey);
    }
    
    authStateChanges.value = false;
  }
  
  
  // Get authenticated HTTP client with token
  Future<http.Client> getAuthenticatedClient() async {
    final token = await getToken();
    final client =  http.Client(); // BrowserClient(); 
    
    if (token != null && token.isValid) {
      return _AuthenticatedClient(client, token.token);
    }
    
    // If token is invalid, try to refresh or return non-authenticated client
    return client;
  }
  
  // Make authenticated request
  // Future<http.Response> authenticatedRequest(
  //   String method,
  //   String endpoint, {
  //   Map<String, String>? headers,
  //   Object? body,
  // }) async {
  //   final token = await getToken();
    
  //   if (token == null || !token.isValid) {
  //     throw Exception('Not authenticated');
  //   }
    
  //   final url = Uri.parse('$baseUrl$endpoint');
  //   final requestHeaders = {
  //     'Content-Type': 'application/json',
  //     'Authorization': 'Bearer ${token.token}',
  //     ...?headers,
  //   };
    
  //   switch (method.toUpperCase()) {
  //     case 'GET':
  //       return http.get(url, headers: requestHeaders);
  //     case 'POST':
  //       return http.post(url, headers: requestHeaders, body: body);
  //     case 'PUT':
  //       return http.put(url, headers: requestHeaders, body: body);
  //     case 'DELETE':
  //       return http.delete(url, headers: requestHeaders);
  //     default:
  //       throw Exception('Unsupported HTTP method: $method');
  //   }
  // }
}

// Custom HTTP client that adds Authorization header
class _AuthenticatedClient extends http.BaseClient {
  final http.Client _inner;
  final String _token;
  
  _AuthenticatedClient(this._inner, this._token);
  
  @override
  Future<http.StreamedResponse> send(http.BaseRequest request) {
    request.headers['Authorization'] = 'Bearer $_token'; 
    // check 여기 hdeader에 credentials: 'include' 넣어야 하나??
    return _inner.send(request);
  }
}
