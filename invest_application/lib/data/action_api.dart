

import 'package:invest_application/presentation/action.dart';
import './auth_api.dart';
import 'dart:convert';
import './config_loader.dart';

abstract class ActionApi {
  Future<List<Event>> getEvents();
  Future<List<AssetInfo>> getAssetInfos();
  Future<bool> recordInvest(int fundId, int assetId, double amount, price);
  Future<bool> runEvent(int eventId);
  Future<bool> toggleEventStatus(int eventId, String status);
}

class ActionApiProvider {
  static ActionApi getApi() {
    if (ConfigLoader.useMock()) {
      return ActionApiHttpMock();
    } else {
      return ActionApiHttp();
    }
  }
}

class ActionApiHttp implements ActionApi {
  final _authService = AuthService();

  @override
  Future<List<Event>> getEvents() async {
  try {

    final url = ConfigLoader.getUrl();
    final client = await _authService.getAuthenticatedClient();
    final response = await client.get(
      Uri.parse('$url/events'),
      headers: {'Content-Type': 'application/json'},
    );
    
    if (response.statusCode == 200) {
      final List<dynamic> eventsJson = jsonDecode(utf8.decode(response.bodyBytes));

      return eventsJson.map((json) => Event.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load events: ${response.statusCode}');
    }
  } catch (e) {
    return []; // Return empty list on error, or you could rethrow the exception
  }
}

  @override
  Future<List<AssetInfo>> getAssetInfos() async {
    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.get(Uri.parse('$url/assets'));
      
      if (response.statusCode == 200) {
        final List<dynamic> jsonList = json.decode(utf8.decode(response.bodyBytes));
        return jsonList.map((json) => AssetInfo(
          id: json['id'].toInt(),
          name: json['name'],
        )).toList();
      } else {
        throw Exception('Failed to load assets: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load assets: $e');
    }
  }
  
  @override
  Future<bool> recordInvest(int fundId, int assetId, double count, price) async {

    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.post(
        Uri.parse('$url/invest'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'fund_id': fundId,
          'asset_id': assetId,
          'count': count,
          'price': price,
        }),
      );

      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }

  @override
  Future<bool> runEvent(int eventId) async {
    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.post(
        Uri.parse('$url/events/launch'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'id': eventId
        })
      );
      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }

  @override
  Future<bool> toggleEventStatus(int eventId, String status) async{

     try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.post(
        Uri.parse('$url/events/switch'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'id': eventId,
          'active': !(status == "active"),
        }),
      );
      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }
}





class ActionApiHttpMock implements ActionApi {

  List<Event> assets = [
      Event(id: 1, title: 'Market Analysis', status: 'Pending', description: 'Analyze market trends for top stocks'),
      Event(id: 2, title: 'Portfolio Rebalancing', status: 'Active', description: 'Rebalance portfolio based on new allocations'),
      Event(id: 3, title: 'Dividend Collection', status: 'Active', description: 'Collect dividends from equity holdings'),
    ];
  
  
  @override
  Future<List<Event>> getEvents() async {
    // Simulate a delay for asynchronous behavior
    await Future.delayed(const Duration(milliseconds: 500));
    // Return a mock list of events
    return assets;
  }

  @override
  Future<List<AssetInfo>> getAssetInfos() async {
    await Future.delayed(const Duration(milliseconds: 500));

    return [
      AssetInfo(id: 1, name: "원"),
      AssetInfo(id: 2, name: "달러"),
      AssetInfo(id: 3, name: "비트코인"),
      AssetInfo(id: 4, name: "금"),
    ];
  }
  
  @override
  Future<bool> recordInvest(int fundId, int assetId, double amount, price) async {

    return true;
  }

  @override
  Future<bool> runEvent(int eventId) async {
    try {      
      // For now, just simulate a successful response
      await Future.delayed(const Duration(seconds: 1));
      return true;
      
    } catch (e) {
      return false;
    }
  }

  @override
  Future<bool> toggleEventStatus(int eventId, String status) async{

    await Future.delayed(const Duration(seconds: 1));

    return true;
  }
}