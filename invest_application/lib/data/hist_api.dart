import 'dart:convert';
import 'package:http/http.dart' as http;
import '../presentation/hist.dart';
import 'package:flutter/material.dart';
import './config_loader.dart';
import 'package:intl/intl.dart';

abstract class HistoryApi {
  Future<List<InvestmentRecord>> getInvestmentHistory(int fundId, DateTimeRange dateRange);
}

class HistoryApiProvider {
  static HistoryApi getApi() {
    if (ConfigLoader.useMock()) {
      return HistoryApiHttpMock();
    } else {
      return HistoryApiHttp();
    }
  }
}

class HistoryApiHttp implements HistoryApi {
  HistoryApiHttp();
  final url = ConfigLoader.getUrl();

  Future<List<InvestmentRecord>> getInvestmentHistory(int fundId, DateTimeRange dateRange) async {

    final response = await http.get(Uri.parse('$url/funds/$fundId/hist?start=${DateFormat('yyyy-MM-dd').format(dateRange.start)}&end=${DateFormat('yyyy-MM-dd').format(dateRange.end)}'));

    if (response.statusCode == 200) {
      final List<dynamic> data = json.decode(response.body);

      return data.map((json) => InvestmentRecord(
        date: DateTime.parse(json['created_at']),
        name: json['asset_name'],
        price: json['price'].toDouble(),
        amount: json['count'] < 0 ? -1*json['count'].toDouble() : json['count'].toDouble(),
        action: json['count'] < 0 ? 'SELL' : 'BUY',
      )).toList();
    } else {
      throw Exception('Failed to load investment history');
    }
  }
}

class HistoryApiHttpMock implements HistoryApi{

  // Simulate API delay
  Future<List<InvestmentRecord>> getInvestmentHistory(
    int fundId, 
    DateTimeRange? dateRange
  ) async {
    await Future.delayed(const Duration(milliseconds: 800));
    
    List<InvestmentRecord> records;
    switch (fundId) {
      case 1:
        records = [
          InvestmentRecord(
            date: DateTime.now().subtract(const Duration(days: 5)),
            name: "AAPL",
            price: 173.50,
            amount: 10,
            action: "BUY",
          ),
          InvestmentRecord(
            date: DateTime.now().subtract(const Duration(days: 3)),
            name: "GOOGL",
            price: 141.80,
            amount: 5,
            action: "BUY",
          ),
          InvestmentRecord(
            date: DateTime.now().subtract(const Duration(days: 1)),
            name: "AAPL",
            price: 178.20,
            amount: 5,
            action: "SELL",
          ),
        ];
        break;
      case 2:
        records = [
          InvestmentRecord(
            date: DateTime.now().subtract(const Duration(days: 2)),
            name: "MSFT",
            price: 338.45,
            amount: 3,
            action: "BUY",
          ),
        ];
        break;
      case 3:
        records = [
          InvestmentRecord(
            date: DateTime.now().subtract(const Duration(days: 1)),
            name: "TSLA",
            price: 238.45,
            amount: 5,
            action: "BUY",
          ),
        ];
        break;
      default:
        return [];
    }
    
    // Filter records by date range if provided
    if (dateRange != null) {
      return records.where((record) =>
        record.date.isAfter(dateRange.start.subtract(const Duration(days: 1))) &&
        record.date.isBefore(dateRange.end.add(const Duration(days: 1)))
      ).toList();
    }
    
    return records;
  }
}
