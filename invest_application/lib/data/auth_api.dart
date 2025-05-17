import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import './config_loader.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
// import 'package:http/browser_client.dart';
// import 'package:shared_preferences/shared_preferences.dart';



// JWT Token model
class AuthToken {
  final String token;
  final DateTime expiryDate;

  AuthToken({required this.token, required this.expiryDate});

  bool get isValid => DateTime.now().isBefore(expiryDate);
}

// abstract class AppStorage {
//   Future<void> write({required String key, required String value});
//   Future<String?> read({required String key});
//   Future<void> delete({required String key});
// }

// class SharedPrefsStorage implements AppStorage {
//   @override
//   Future<void> write({required String key, required String value}) async {
//     final prefs = await SharedPreferences.getInstance();
//     await prefs.setString(key, value);
//   }

//   @override
//   Future<String?> read({required String key}) async {
//     final prefs = await SharedPreferences.getInstance();
//     return prefs.getString(key);
//   }

//   @override
//   Future<void> delete({required String key}) async {
//     final prefs = await SharedPreferences.getInstance();
//     await prefs.remove(key);
//   }
// }

class AuthService {

  final String baseUrl = ConfigLoader.getUrl();
  final FlutterSecureStorage _storage = const FlutterSecureStorage(); // in-device storage
  // final AppStorage _storage =  SharedPrefsStorage();
  
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
  
  // Login and get JWT token
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
        
        // Save to secure storage
        await _saveToken(token, expiryDate);
        // Notify UI about auth state change
        authStateChanges.value = true;
        return true;
      } else {
        return false;
      }
    } catch (e) {
      return false;
    }
  }
  
  // Save token to secure storage
  Future<void> _saveToken(String token, DateTime expiryDate) async {
    await _storage.write(key: _tokenKey, value: token);
    await _storage.write(key: _tokenExpiryKey, value: expiryDate.millisecondsSinceEpoch.toString());
  }
  
  // Get token from secure storage
  Future<AuthToken?> getToken() async {
    final token = await _storage.read(key: _tokenKey);
    final expiryString = await _storage.read(key: _tokenExpiryKey);
    
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
    await _storage.delete(key: _tokenKey);
    await _storage.delete(key: _tokenExpiryKey);
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
