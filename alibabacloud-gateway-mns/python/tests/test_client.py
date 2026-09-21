# -*- coding: utf-8 -*-
import re
import unittest
from datetime import timezone
from email.utils import parsedate_to_datetime

from alibabacloud_gateway_mns.client import Client
from alibabacloud_tea_util.client import Client as UtilClient


class TestClient(unittest.TestCase):

    def _new_client(self):
        # Skip generated __init__: it currently references an undefined name and
        # is not required for the date conversion helpers under test.
        return Client.__new__(Client)

    def test_get_date_iso8601_from_utc(self):
        client = self._new_client()
        cases = (
            ("Thu, 06 Feb 2020 07:32:54 GMT", "20200206T073254Z"),
            ("Tue, 03 Jun 2025 11:23:48 GMT", "20250603T112348Z"),
            ("Wed, 03 Dec 2025 11:23:48 GMT", "20251203T112348Z"),
            ("Thu, 01 Jan 2026 00:00:00 GMT", "20260101T000000Z"),
            ("Thu, 31 Dec 2026 23:59:59 GMT", "20261231T235959Z"),
        )
        for rfc, expected in cases:
            self.assertEqual(expected, client._get_date_iso8601_from_utc(rfc))

    def test_get_date_iso8601_from_utc_round_trip(self):
        client = self._new_client()
        rfc = UtilClient.get_date_utcstring()
        self.assertRegex(
            rfc,
            r"^[A-Z][a-z]{2}, \d{2} [A-Z][a-z]{2} \d{4} \d{2}:\d{2}:\d{2} GMT$",
        )
        parsed = parsedate_to_datetime(rfc).astimezone(timezone.utc)
        expected = parsed.strftime("%Y%m%dT%H%M%SZ")
        actual = client._get_date_iso8601_from_utc(rfc)
        self.assertEqual(expected, actual)
        self.assertRegex(actual, r"^\d{8}T\d{6}Z$")
