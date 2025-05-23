import './config_loader.dart';
import './auth_api.dart';
import 'dart:convert';
import 'package:flutter/material.dart';


abstract class MarketApi {
  Future<bool> updateMarketStatus(MarketStatus status);
  Future<MarketStatus> getMarketStatus();
  Future<Map<String, dynamic>> getIndexs();
}

class MarketApiProvider {
  static MarketApi getApi() {
    if (ConfigLoader.useMock()) {
      return MarketApiHttpMock();
    } else {
      return MarketApiHttp();
    }
  }
}

class MarketApiHttp implements MarketApi {
  final _authService = AuthService();

  @override
  Future<MarketStatus> getMarketStatus() async {
    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.get(Uri.parse('$url/market'));
      
      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        int status = data['Status']; // This will be an int
        switch(status){
          case 1:
            return MarketStatus.MAJOR_BEAR;
          case 2:
            return MarketStatus.BEAR;
          case 3:
            return MarketStatus.VOLATILE;
          case 4:
            return  MarketStatus.BULL;
          case 5:
            return MarketStatus.MAJOR_BULL;
          default:
          throw Exception('Invalid status');
    }

      } else {
        throw Exception('Failed to load assets: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load assets: $e');
    }
  }
  
  @override
  Future<bool> updateMarketStatus(MarketStatus status) async{

    int value = 0;
    switch(status){
      case MarketStatus.MAJOR_BEAR:
        value = 1;
      case MarketStatus.BEAR:
        value = 2;
      case MarketStatus.VOLATILE:
        value = 3;
      case MarketStatus.BULL:
        value = 4;
      case MarketStatus.MAJOR_BULL:
        value = 5;
    }

    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.post(
        Uri.parse('$url/market'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'status': value
        }),
      );
      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }

  @override
  Future<Map<String, dynamic>> getIndexs() async {
   try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.get(Uri.parse('$url/market/weekly_indicators'));
      if (response.statusCode == 200) {
        Map<String, dynamic> data = json.decode(response.body);
        return data;
        
        // final fearGreedWeek = List<int>.from(data['fear_greed']);
        // final nasdaqWeek = List<double>.from(data['nasdaq'].map((e) => (e as num).toDouble()));
        // final sp500Week = List<double>.from(data['sp500'].map((e) => (e as num).toDouble()));
        // return {     
        //   'Fear & Greed Index': {
        //     'value': fearGreedWeek[6].toString(),
        //     'status': fearGreedWeek[6] > 50? 'GREED' : 'FEAR',
        //     'graph': fearGreedWeek
        //   },
        //   'NASDAQ' : {
        //     'value': nasdaqWeek[6].toString(),
        //     'status': (100 * ((nasdaqWeek[6]-nasdaqWeek[5]) / nasdaqWeek[6])).toStringAsFixed(2)+'%' ,
        //     'graph': nasdaqWeek
        //   },
        //   'S&P 500' : {
        //     'value': sp500Week[6].toString(),
        //     'status': (100 * ((sp500Week[6]-sp500Week[5]) / sp500Week[6])).toStringAsFixed(2)+'%' ,
        //     'graph': sp500Week
        //   }
        // };
        
      } else {
        throw Exception('Failed to load index: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load index: $e');
    }
  }
}

class MarketApiHttpMock implements MarketApi{

  @override
  Future<MarketStatus> getMarketStatus() async{
    return MarketStatus.BEAR;
  }
  
  @override
  Future<bool> updateMarketStatus(MarketStatus status) async{
    return true;
  }

  @override
  Future<Map<String, dynamic>> getIndexs() async{

    await Future.delayed(const Duration(milliseconds: 800));
    return {
      'fearGreed': {
        'value': 75,
        'status': 'Greed',
        'weeklyData': [65, 70, 72, 68, 73, 75, 75], // Mock weekly data
      },
      'nasdaq' : {
        'value': 15055.65,
        'status': '0.85%',
        'weeklyData': [14800, 14900, 15100, 14950, 15000, 15055.65], // Mock weekly data
      },
      'sp500' : {
        'value': 15055.65,
        'status': '0.85%',
        'weeklyData': [14800, 14900, 15100, 14950, 15000, 15055.65], // Mock weekly data
    }
    };
  }
  
}

enum MarketStatus {
  MAJOR_BEAR,
  BEAR,
  VOLATILE,
  BULL,
  MAJOR_BULL
}

// Extension to get display names and colors for market status
extension MarketStatusExtension on MarketStatus {
  String get displayName {
    switch (this) {
      case MarketStatus.MAJOR_BEAR:
        return 'Major Bear';
      case MarketStatus.BEAR:
        return 'Bear';
      case MarketStatus.VOLATILE:
        return 'Volatile';
      case MarketStatus.BULL:
        return 'Bull';
      case MarketStatus.MAJOR_BULL:
        return 'Major Bull';
    }
  }

  Color get color {
    switch (this) {
      case MarketStatus.MAJOR_BEAR:
        return Colors.red[900]!;
      case MarketStatus.BEAR:
        return Colors.red[400]!;
      case MarketStatus.VOLATILE:
        return Colors.amber[600]!;
      case MarketStatus.BULL:
        return Colors.green[400]!;
      case MarketStatus.MAJOR_BULL:
        return Colors.green[900]!;
    }
  }
}