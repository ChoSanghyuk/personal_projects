import './config_loader.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';


abstract class MarketApi {
  Future<Map<String, dynamic>> getIndexs();
}

class MarketApiHttp implements MarketApi {
  
  @override
  Future<Map<String, dynamic>> getIndexs() async {
   try {
      final url = ConfigLoader.getUrl();
      final response = await http.get(Uri.parse('$url/market/weekly_indicators'));
      if (response.statusCode == 200) {
        final data = json.decode(response.body);

        final fearGreedWeek = List<int>.from(data['fear_greed']);
        final nasdaqWeek = List<double>.from(data['nasdaq'].map((e) => (e as num).toDouble()));
        final sp500Week = List<double>.from(data['sp500'].map((e) => (e as num).toDouble()));

        return {     
          'fearGreed': {
            'value': fearGreedWeek[6],
            'status': fearGreedWeek[6] > 50? 'GREED' : 'FEAR',
            'weeklyData': fearGreedWeek
          },
          'nasdaq' : {
            'value': nasdaqWeek[6],
            'change': 100 * ((nasdaqWeek[6]-nasdaqWeek[5]) / nasdaqWeek[6]) ,
            'weeklyData': nasdaqWeek
          },
          'sp500' : {
            'value': sp500Week[6],
            'change': 100 * ((sp500Week[6]-sp500Week[5]) / sp500Week[6]) ,
            'weeklyData': sp500Week
          }
        };
      } else {
        throw Exception('Failed to load index: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load index: $e');
    }
  }

  // Future<Map<String, dynamic>> getNasdaqData() async {
  //   await Future.delayed(const Duration(milliseconds: 800));
  //   return {
  //     'value': 15055.65,
  //     'change': 0.85,
  //     'weeklyData': [14800, 14900, 15100, 14950, 15000, 15055.65], // Mock weekly data
  //   };
  // }

  // Future<Map<String, dynamic>> getSP500Data() async {
  //   await Future.delayed(const Duration(milliseconds: 800));
  //   return {
  //     'value': 4783.35,
  //     'change': 1.2,
  //   };
  // }
}
