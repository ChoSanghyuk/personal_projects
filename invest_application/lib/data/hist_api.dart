import 'dart:convert';
import 'package:http/http.dart' as http;
import '../presentation/hist.dart';
import 'package:flutter/material.dart';
import './config_loader.dart';
import 'package:intl/intl.dart';

abstract class HistoryApi {
  Future<List<InvestmentRecord>> getInvestmentHistory(int fundId, DateTimeRange dateRange);
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
