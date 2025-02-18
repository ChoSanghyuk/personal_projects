import 'package:flutter/material.dart';
import '../presentation/hist.dart';

class HistoryApiMock {
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
            action: "buy",
          ),
          InvestmentRecord(
            date: DateTime.now().subtract(const Duration(days: 3)),
            name: "GOOGL",
            price: 141.80,
            amount: 5,
            action: "buy",
          ),
          InvestmentRecord(
            date: DateTime.now().subtract(const Duration(days: 1)),
            name: "AAPL",
            price: 178.20,
            amount: 5,
            action: "sell",
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
            action: "buy",
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
            action: "buy",
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
